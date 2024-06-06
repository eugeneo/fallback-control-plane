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
	envoy "github.com/eugeneo/fallback-control-plane/envoy"
)

var (
	l        envoy.Logger
	port     uint
	nodeID   string
	upstream = flag.String("upstream", "localhost:3000", "upstream server")
)

func init() {
	l = envoy.Logger{}

	flag.BoolVar(&l.Debug, "debug", false, "Enable xDS server debug logging")

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")
}

func main() {
	flag.Parse()

	// Create a cache
	cache := cache.NewSnapshotCache(false, cache.IDHash{}, l)
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
	snapshot := envoy.GenerateSnapshot((*upstream)[:sep], uint32(upstreamPort))
	if err := snapshot.Consistent(); err != nil {
		l.Errorf("snapshot inconsistency: %+v\n%+v", snapshot, err)
		os.Exit(1)
	}
	l.Debugf("will serve snapshot %+v", snapshot)

	// Add the snapshot to the cache
	if err := cache.SetSnapshot(context.Background(), nodeID, snapshot); err != nil {
		l.Errorf("snapshot error %q for %+v", err, snapshot)
		os.Exit(1)
	}

	// Run the xDS server
	ctx := context.Background()
	cb := &test.Callbacks{Debug: l.Debug}
	srv := server.NewServer(ctx, cache, cb)
	envoy.RunServer(srv, port)
}
