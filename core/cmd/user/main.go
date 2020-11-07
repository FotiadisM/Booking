package main

import (
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
		Addr:         ":8080",
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
		logger.Log("msg", "HTTP", "addr", ":8080")
		errs <- s.ListenAndServe()
	}()

	logger.Log("exit", <-errs)
}
