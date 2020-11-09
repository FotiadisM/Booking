package searchconsumer

import (
	"net/url"

	"github.com/go-kit/kit/transport/http"
)

// AddListingHandler returns a http Handler
func AddListingHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeAddListingEndpoint(svc),
		decodeAddListingRequest,
		http.EncodeJSONResponse,
	)
}

// AddReviewHandler returns a http Handler
func AddReviewHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeAddReviewEndpoint(svc),
		decodeAddReviewRequest,
		http.EncodeJSONResponse,
	)
}

// AddListingClient returns a http Client
func AddListingClient(u *url.URL) *http.Client {
	return http.NewClient(
		"POST",
		u,
		http.EncodeJSONRequest,
		decodeAddListingResponse,
	)
}
