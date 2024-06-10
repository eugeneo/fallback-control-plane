package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/v3"
	example "github.com/eugeneo/fallback-control-plane/envoy"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	l        example.Logger
	port     uint
	nodeID   string
	upstream = flag.String("upstream", "localhost:3000", "upstream server")
)

func init() {
	l = example.Logger{}

	flag.BoolVar(&l.Debug, "debug", false, "Enable xDS server debug logging")

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")
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
	if err := cache.SetSnapshot(context.Background(), nodeID, snapshot); err != nil {
		l.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
	cb := &test.Callbacks{Debug: l.Debug}
	srv := server.NewServer(ctx, cache, cb)
	example.RunServer(srv, port)
}
