package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	searchconsumer "github.com/FotiadisM/booking/core/services/search_consumer"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func main() {
	host := flag.String("host", "localhost", "http host")
	port := flag.String("port", "8060", "http port")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "service", "search_consumer")

	repo := newRepository()

	var svc searchconsumer.ServiceModel
	svc = searchconsumer.NewService(repo)
	svc = searchconsumer.LoggingMiddleware{Logger: logger, Next: svc}

	r := mux.NewRouter()
	r.Handle("/search_consumer/listing", searchconsumer.AddListingHandler(svc)).Methods("POST")
	r.Handle("/search_consumer/review", searchconsumer.AddReviewHandler(svc)).Methods("POST")

	s := http.Server{
		Addr:         *host + ":" + *port,
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("msg", "HTTP", "host", *host, "port", *port)
		errs <- s.ListenAndServe()
	}()

	logger.Log("exit", <-errs)
}
