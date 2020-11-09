package review

import "github.com/go-kit/kit/transport/http"

// GetAllHandler returns a http Handler
func GetAllHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeGetAllEndpoint(svc),
		decodeGetAllRequest,
		http.EncodeJSONResponse,
	)
}

// GetByListingIDHandler returns a http Handler
func GetByListingIDHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeGetByListingIDEndpoint(svc),
		decodeGetByListingIDRequest,
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
