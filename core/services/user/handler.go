package user

import "github.com/go-kit/kit/transport/http"

// GetByIDHandler returns a http Handler
func GetByIDHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeGetByIDEndpoint(svc),
		decodeGetByIDRequest,
		http.EncodeJSONResponse,
	)
}

// CreateHandler returns a http Handler
func CreateHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		http.EncodeJSONResponse,
	)
}
