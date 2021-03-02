package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vivian-hua/civic-qa/services/common/config"
	aggregator "github.com/vivian-hua/civic-qa/services/logAggregator/pkg/middleware"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.0"
)

func main() {
	// config
	var cfg config.Provider = &config.EnvProvider{}
	cfg.SetVerbose(true)

	// routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// middleware
	router.Use(aggregator.NewAggregatorMiddleware(&aggregator.Config{
		AggregatorAddress: cfg.GetOrFallback("AGG_ADDR", "http://localhost:8888"),
		ServiceName:       "form",
		StdoutErrors:      true,
		Timeout:           10 * time.Second,
	}))

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello world!")
	})

	// start server
	addr := cfg.GetOrFallback("ADDR", ":7070")
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
