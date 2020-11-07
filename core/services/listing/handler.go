package listing

import "github.com/go-kit/kit/transport/http"

// GetAllHandler returns a http Handler
func GetAllHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeGetAllEndpoint(svc),
		decodeGetAllRequest,
		http.EncodeJSONResponse,
	)
}

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
