package main

import (
	"net/http"
	"os"

	"github.com/FotiadisM/booking/core/services/user"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc user.ServiceModel
	svc = user.Service{}
	svc = user.LoggingMiddleware{Logger: logger, Next: svc}

	getUserHandler := httptransport.NewServer(
		user.MakeGetUserEndpoint(svc),
		user.DecodeGetUserRequest,
		httptransport.EncodeJSONResponse,
	)

	createUserHandler := httptransport.NewServer(
		user.MakeCreateUserEndpoint(svc),
		user.DecodeCreateUserRequest,
		httptransport.EncodeJSONResponse,
	)

	http.Handle("/getuser", getUserHandler)
	http.Handle("/createUser", createUserHandler)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
