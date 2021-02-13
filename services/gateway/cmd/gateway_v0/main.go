package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
)

func main() {

	// Variables
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	// Routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// Routes
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello world!")
	})

	// Middleware
	router.Use(middleware.NewLoggingMiddleware(LoggingOutput))

	// Start
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))

}
