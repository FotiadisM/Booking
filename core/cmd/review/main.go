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

	"github.com/FotiadisM/booking/core/services/review"
	searchconsumer "github.com/FotiadisM/booking/core/services/search_consumer"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("port", "8080", "http port")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "service", "review")

	repo := newRepository()

	var svc review.ServiceModel

	scu := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/search_consumer/review",
	}
	sccl := searchconsumer.AddReviewClient(scu)

	lu := &url.URL{
		Scheme: "http",
		Host:   "localhost:8090",
		Path:   "/listing",
	}
	lcl := searchconsumer.AddReviewClient(lu)

	svc = review.NewService(repo, sccl.Endpoint(), lcl.Endpoint())
	svc = review.LoggingMiddleware{Logger: logger, Next: svc}

	r := mux.NewRouter()
	r.Handle("/reviews/{id}", review.GetByListingIDHandler(svc)).Methods("GET")
	r.Handle("/reviews", review.GetAllHandler(svc)).Methods("GET")
	r.Handle("/reviews", review.CreateHandler(svc)).Methods("POST")

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
