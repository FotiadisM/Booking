package main

import (
	"context"
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
	host := flag.String("host", "localhost", "http host")
	port := flag.String("port", "8090", "http port")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "service", "listing")

	repo := newRepository()

	var svc listing.ServiceModel
	svc = listing.NewService(repo)
	svc = listing.LoggingMiddleware{Logger: logger, Next: svc}

	r := mux.NewRouter()
	r.Handle("/listings/{id}", listing.GetByIDHandler(svc)).Methods("GET")
	r.Handle("/listings", listing.GetAllHandler(svc)).Methods("GET")
	r.Handle("/listings", listing.CreateHandler(svc)).Methods("POST")

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

	type addListingRequest struct {
		L *listing.Listing `json:"listing"`
	}

	type addListingResponse struct {
		Err string `json:"error,omitempty"`
	}

	u := &url.URL{
		Scheme: "http",
		Host:   "localhost:8060",
		Path:   "/search_consumer/listing",
	}

	cl := searchconsumer.AddListingClient(u)
	e := cl.Endpoint()

	req := addListingRequest{L: &listing.Listing{UserID: "1234"}}

	_, err := e(context.Background(), req)
	if err != nil {
		logger.Log("error", err)
	}

	logger.Log("exit", <-errs)
}
