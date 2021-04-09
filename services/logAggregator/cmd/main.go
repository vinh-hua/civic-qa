package main

import (
	// standard
	"log"
	"net/http"

	// 3rd party
	"github.com/gorilla/mux"

	// common
	"github.com/team-ravl/civic-qa/services/common/config"

	// internal
	"github.com/team-ravl/civic-qa/services/logAggregator/internal/context"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.1"
)

func main() {

	// config
	var cfg config.Provider = &config.EnvProvider{}
	cfg.SetVerbose(true)

	// Routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// Handler context
	ctx, err := context.BuildContext(cfg)
	if err != nil {
		log.Fatalf("Failed to create handler context: %v", err)
	}

	// Routes
	api.HandleFunc("/log", ctx.HandleLog)
	api.HandleFunc("/query", ctx.HandleQuery)

	// Start Server
	addr := cfg.GetOrFallback("ADDR", ":8888")
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
