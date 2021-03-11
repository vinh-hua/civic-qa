package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/team-ravl/civic-qa/services/common/config"
	aggregator "github.com/team-ravl/civic-qa/services/logAggregator/pkg/middleware"
	"github.com/team-ravl/civic-qa/services/mailto/internal/context"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.0"
)

func main() {
	var cfg config.Provider = &config.EnvProvider{}
	cfg.SetVerbose(true)

	// routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// middleware
	router.Use(aggregator.NewAggregatorMiddleware(&aggregator.Config{
		AggregatorAddress: cfg.GetOrFallback("AGG_ADDR", "http://localhost:8888"),
		ServiceName:       "mailto",
		SkipSuccesses:     true,
		StdoutErrors:      true,
		Timeout:           10 * time.Second,
	}))

	// request context
	ctx, err := context.BuildContext(cfg)
	if err != nil {
		log.Fatalf("Failed to create handler context: %v", err)
	}

	// routes
	api.HandleFunc("/mailto", ctx.Mailto).Methods("POST")

	// start server
	addr := cfg.GetOrFallback("ADDR", ":6060")
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
