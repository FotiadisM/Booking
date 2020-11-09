package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FotiadisM/booking/core/services/review"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func main() {
	host := flag.String("host", "localhost", "http host")
	port := flag.String("port", "8070", "http port")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "service", "review")

	repo := newRepository()

	var svc review.ServiceModel
	svc = review.NewService(repo)
	svc = review.LoggingMiddleware{Logger: logger, Next: svc}

	r := mux.NewRouter()
	r.Handle("/reviews/{id}", review.GetByListingIDHandler(svc)).Methods("GET")
	r.Handle("/reviews", review.GetAllHandler(svc)).Methods("GET")
	r.Handle("/reviews", review.CreateHandler(svc)).Methods("POST")

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
