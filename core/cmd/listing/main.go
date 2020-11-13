package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FotiadisM/booking/core/services/listing"
	searchconsumer "github.com/FotiadisM/booking/core/services/search_consumer"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("port", "8080", "http port")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "service", "listing")

	repo := newRepository()

	var svc listing.ServiceModel

	u := &url.URL{
		Scheme: "http",
		Host:   "localhost:8060",
		Path:   "/search_consumer/listing",
	}

	cl := searchconsumer.AddListingClient(u)
	e := cl.Endpoint()

	svc = listing.NewService(repo, e)
	svc = listing.LoggingMiddleware{Logger: logger, Next: svc}

	r := mux.NewRouter()
	r.Handle("/listings/{id}", listing.GetByIDHandler(svc)).Methods("GET")
	r.Handle("/listings", listing.GetAllHandler(svc)).Methods("GET")
	r.Handle("/listings", listing.CreateHandler(svc)).Methods("POST")
	r.Handle("/listings", listing.AddReviewToListingHandler(svc)).Methods("PUT")

	s := http.Server{
		Addr:         ":" + *port,
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
		logger.Log("msg", "HTTP", "port", *port)
		errs <- s.ListenAndServe()
	}()

	logger.Log("exit", <-errs)
}
