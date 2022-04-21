package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/edfoh/go-api-starter/cmd/bootstrap"
	"github.com/edfoh/go-api-starter/internal/adding"
	http2 "github.com/edfoh/go-api-starter/internal/http"
	"github.com/edfoh/go-api-starter/pkg/log"
)

var (
	logLevelVar    = flag.String("logLevel", "error", "Set log level. Options: debug, warn, info, error, fatal, panic. Default is error")
	storageTypeVar = flag.String("storageType", "memory", "Set storage type. Options: memory, database. Default is memory")
)

func main() {
	flag.Parse()
	logLevel := resolveLogLevel(*logLevelVar)

	if err := run(logLevel, *storageTypeVar); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run(level log.Level, storageType string) error {
	dep := bootstrap.Dependencies{}
	logger := dep.Logger(level)
	router := dep.Router()

	var serviceAdder adding.Service
	switch storageType {
	case "memory":
		storage := dep.StorageMemory(logger)
		serviceAdder = dep.ServiceAdder(storage, logger)
	}

	server := http2.NewServer(logger, router, serviceAdder)
	return http.ListenAndServe(":8080", server)
}

func resolveLogLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "debug":
		return log.DebugLevel
	case "warn":
		return log.WarnLevel
	case "info":
		return log.InfoLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	}
	return log.ErrorLevel
}
