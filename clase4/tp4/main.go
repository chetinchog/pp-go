package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"time"

	admissioncontrol "github.com/elithrar/admission-control"
	log "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

const port = "5000"
const host = "localhost"

func main() {
	r := mux.NewRouter()
	PreciosRouter(r)

	var logger log.Logger
	// Logfmt is a structured, key=val logging format that is easy to read and parse
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	// Direct any attempts to use Go's log package to our structured logger
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// Log the timestamp (in UTC) and the callsite (file + line number) of the logging
	// call for debugging in the future.
	// logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	// Create an instance of our LoggingMiddleware with our configured logger
	loggedRouter := admissioncontrol.LoggingMiddleware(logger)(r)

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         host + ":" + port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	fmt.Println("Server listening in " + host + ":" + port)
	stdlog.Fatal(srv.ListenAndServe())
}
