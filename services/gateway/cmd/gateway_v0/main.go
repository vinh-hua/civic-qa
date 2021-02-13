package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/vivian-hua/civic-qa/services/common"
)

const (
	VersionBase = "v0"
	APIVersion  = "0.0.0"
)

var (
	LoggingOutput = os.Stdout
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix(VersionBase)

	router.Use(common.NewLoggingMiddleware(LoggingOutput))

}
