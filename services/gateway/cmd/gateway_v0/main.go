package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vivian-hua/civic-qa/services/common/environment"
	"github.com/vivian-hua/civic-qa/services/common/middleware"
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
	addr = environment.GetEnvOrFallback("ADDR", ":80")
)

func main() {

	// Routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// Middleware
	router.Use(middleware.NewLoggingMiddleware(LoggingOutput))

	// Routes
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello world!")
	})

	// Start Server
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))

}
