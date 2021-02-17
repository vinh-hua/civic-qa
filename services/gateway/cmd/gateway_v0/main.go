package main

import (
	// standard
	"io"
	"log"
	"net/http"
	"os"

	// 3rd party
	"github.com/gorilla/mux"

	// common
	"github.com/vivian-hua/civic-qa/services/common/environment"
	commonMiddleware "github.com/vivian-hua/civic-qa/services/common/middleware"
	logClient "github.com/vivian-hua/civic-qa/services/logAggregator/pkg/client"

	// internal
	"github.com/vivian-hua/civic-qa/services/gateway/internal/middleware"
)

const (
	// VersionBase is the API route base
	VersionBase = "/v0"
	// APIVersion is the API semantic version
	APIVersion = "v0.0.0"
)

var (
	// LoggingOutput is a file that recieves log outputs
	LoggingOutput = os.Stdout

	// Environment
	addr           = environment.GetEnvOrFallback("ADDR", ":80")
	aggregatorAddr = environment.GetEnvOrFallback("AGG_ADDR", "http://localhost:8888")
)

func main() {

	// Routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// Middleware
	router.Use(middleware.NewCorrelatorMiddleware)
	router.Use(commonMiddleware.NewLoggingMiddleware(LoggingOutput))

	// Handlers and clients
	aggregator, err := logClient.NewLogClient(aggregatorAddr)
	if err != nil {
		log.Fatalf("Failed to create aggregator client: %v", err)
	}

	// Routes
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		aggregator.Log(r.Header.Get("X-Correlation-ID"), "Gateway", 200, "Hello!")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello world!")
	})

	// Start Server
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
