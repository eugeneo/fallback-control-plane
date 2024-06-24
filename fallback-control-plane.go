package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	cs "github.com/eugeneo/fallback-control-plane/grpc/interop/grpc_testing/xdsconfig"
	"github.com/eugeneo/fallback-control-plane/testcontrolplane"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	example "github.com/eugeneo/fallback-control-plane/envoy"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	l        example.Logger
	port     = flag.Uint("port", 3333, "Port to listen on")
	nodeID   = flag.String("node", "test-id", "Node ID")
	upstream = flag.String("upstream", "localhost:3000", "upstream server")
)

func init() {
	l = example.Logger{}
	flag.BoolVar(&l.Debug, "debug", false, "Enable xDS server debug logging")
}

type ControlService struct {
	cs.UnsafeXdsConfigControlServiceServer
	Cb *testcontrolplane.Callbacks
}

func (srv *ControlService) StopOnRequest(_ context.Context, req *cs.StopOnRequestRequest) (*cs.StopOnRequestResponse, error) {
	srv.Cb.AddFilter(req.GetResourceType(), req.GetResourceName())
	res := cs.StopOnRequestResponse{}
	for _, f := range srv.Cb.GetFilters() {
		res.Filters = append(res.Filters, &cs.StopOnRequestResponse_ResourceFilter{ResourceType: f.ResourceType, ResourceName: f.ResourceName})
	}
	return &res, nil
}

func main() {
	flag.Parse()
	sep := strings.LastIndex(*upstream, ":")
	if sep < 0 {
		l.Errorf("Incorrect upstream host name: %+v", upstream)
		os.Exit(1)
	}
	upstreamPort, err := strconv.Atoi((*upstream)[sep+1:])
	if err != nil {
		l.Errorf("Bad upstream port: %+v\n%+v", (*upstream)[sep+1:], err)
		os.Exit(1)
	}

	// Create the snapshot that we'll serve to Envoy
	snapshot := example.GenerateSnapshot((*upstream)[:sep], uint32(upstreamPort))
	if err := snapshot.Consistent(); err != nil {
		l.Errorf("snapshot inconsistency: %+v", err)
		for _, r := range snapshot.Resources {
			for name, resource := range r.Items {
				bytes, err := prototext.MarshalOptions{Multiline: true}.Marshal(resource.Resource)
				if err != nil {
					l.Errorf("Can't marshal %s", name)
				} else {
					l.Errorf("Resource: %s\n%s",
						resource.Resource,
						string(bytes))
				}
			}
		}
		os.Exit(1)
	}
	l.Debugf("will serve snapshot:")
	for _, values := range snapshot.Resources {
		for name, item := range values.Items {
			text, err := prototext.MarshalOptions{Multiline: true}.Marshal(item.Resource)
			if err != nil {
				l.Errorf("Resource %+v, error: %+v", name, err)
			} else {
				l.Debugf("%+v => %+v", name, string(text))
			}
		}
	}

	// Create a cache
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, l)
	// Add the snapshot to the cache
	if err := cache.SetSnapshot(context.Background(), *nodeID, snapshot); err != nil {
		l.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
	cb := &testcontrolplane.Callbacks{Debug: l.Debug, Filters: make(map[string]map[string]bool)}
	controlService := &ControlService{Cb: cb}
	srv := server.NewServer(ctx, cache, cb)
	example.RunServer(srv, controlService, *port)
}
