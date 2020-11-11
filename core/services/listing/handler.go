package listing

import (
	"net/url"

	"github.com/go-kit/kit/transport/http"
)

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

// AddReviewToListingHandler returns a http Handler
func AddReviewToListingHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeAddReviewToListing(svc),
		decodeAddReviewToListingRequest,
		http.EncodeJSONResponse,
	)
}

// AddReviewToListingClient returns a http Handler
func AddReviewToListingClient(u *url.URL) *http.Client {
	return http.NewClient(
		"POST",
		u,
		http.EncodeJSONRequest,
		decodeAddReviewToListingRespnose,
	)
}
