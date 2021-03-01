package main

import (
	// standard
	"io"
	"log"
	"net/http"
	"os"
	"time"

	// 3rd party
	"github.com/gorilla/mux"

	// common
	"github.com/vivian-hua/civic-qa/services/common/config"
	commonMiddleware "github.com/vivian-hua/civic-qa/services/common/middleware"

	//pkg
	aggregator "github.com/vivian-hua/civic-qa/services/logAggregator/pkg/middleware"

	// internal
	"github.com/vivian-hua/civic-qa/services/gateway/internal/middleware"
	"github.com/vivian-hua/civic-qa/services/gateway/internal/proxy"
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
	// config
	var cfg config.Provider = &config.EnvProvider{}
	cfg.SetVerbose(true)

	// Routers
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase).Subrouter()

	// Middleware
	router.Use(middleware.NewCorrelatorMiddleware)
	router.Use(commonMiddleware.NewLoggingMiddleware(LoggingOutput))
	router.Use(aggregator.NewAggregatorMiddleware(&aggregator.Config{
		AggregatorAddress: cfg.GetOrFallback("AGG_ADDR", "http://localhost:8888"),
		ServiceName:       "gateway",
		StdoutErrors:      true,
		Timeout:           10 * time.Second,
	}))

	// Routes
	accountService := cfg.GetOrFallback("ACCOUNT_SVC", "http://localhost:8080/v0")
	api.Handle("/signup", proxy.NewProxy(proxy.MustParse(accountService+"/signup")))
	api.Handle("/login", proxy.NewProxy(proxy.MustParse(accountService+"/login")))
	api.Handle("/logout", proxy.NewProxy(proxy.MustParse(accountService+"/logout")))
	api.Handle("/getsession", proxy.NewProxy(proxy.MustParse(accountService+"/getsession")))

	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello world!")
	})

	// Start Server
	addr := cfg.GetOrFallback("ADDR", ":80")
	log.Printf("Server %s running on %s", APIVersion, addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
