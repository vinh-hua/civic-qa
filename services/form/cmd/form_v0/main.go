package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/team-ravl/civic-qa/services/common/config"
	"github.com/team-ravl/civic-qa/services/form/internal/context"
	aggregator "github.com/team-ravl/civic-qa/services/logAggregator/pkg/middleware"
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
	api.HandleFunc("/forms", ctx.HandleGetForms).Methods("GET")
	api.HandleFunc("/forms", ctx.HandleCreateForm).Methods("POST")
	api.HandleFunc("/forms/{formID:[0-9]+}", ctx.HandleGetSpecificForm).Methods("GET")

	api.HandleFunc("/responses", ctx.HandleGetResponses).Methods("GET")
	api.HandleFunc("/responses/{responseID:[0-9]+}", ctx.HandlePatchResponse).Methods("PATCH")
	api.HandleFunc("/responses/{responseID:[0-9]+}", ctx.HandleGetSpecificResponse).Methods("GET")
	api.HandleFunc("/responses/{responseID:[0-9]+}/tags", ctx.HandleGetTags).Methods("GET")
	api.HandleFunc("/responses/{responseID:[0-9]+}/tags", ctx.HandlePostTag).Methods("POST")
	api.HandleFunc("/responses/{responseID:[0-9]+}/tags", ctx.HandleDeleteTag).Methods("DELETE")

	api.HandleFunc("/form/{formID:[0-9]+}", ctx.HandleGetForm).Methods("GET")
	api.HandleFunc("/form/{formID:[0-9]+}", ctx.HandlePostForm).Methods("POST")

	api.HandleFunc("/tags", ctx.HandleGetAllTags).Methods("GET")

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello world!")
	})

	// start server
	addr := cfg.GetOrFallback("ADDR", ":7070")
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
