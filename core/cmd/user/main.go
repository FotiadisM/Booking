package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FotiadisM/booking/core/services/user"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.String("port", "8080", "http port")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "service", "user")

	repo := newRepository()

	var svc user.ServiceModel
	svc = user.NewService(repo)
	svc = user.LoggingMiddleware{Logger: logger, Next: svc}

	r := mux.NewRouter()
	r.Handle("/users/{id}", user.GetByIDHandler(svc)).Methods("GET")
	r.Handle("/users", user.CreateHandler(svc)).Methods("POST")

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
