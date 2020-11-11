package searchconsumer

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FotiadisM/booking/core/services/listing"
	"github.com/FotiadisM/booking/core/services/review"
)

type addListingRequest struct {
	L *listing.Listing `json:"listing"`
}

type addListingResponse struct {
	Err string `json:"error,omitempty"`
}

func decodeAddListingRequest(_ context.Context, r *http.Request) (interface{}, error) {

	l := &listing.Listing{}
	if err := json.NewDecoder(r.Body).Decode(l); err != nil {
		return nil, err
	}

	return addListingRequest{L: l}, nil
}

func decodeAddListingResponse(_ context.Context, r *http.Response) (response interface{}, err error) {

	res := &addListingResponse{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err
	}

	return response, nil
}

type addReviewRequest struct {
	R *review.Review `json:"review"`
}

type addReviewResponse struct {
	Err string `json:"error,omitempty"`
}

func decodeAddReviewRequest(_ context.Context, r *http.Request) (interface{}, error) {

	rev := &review.Review{}
	if err := json.NewDecoder(r.Body).Decode(rev); err != nil {
		return nil, err
	}

	return addReviewRequest{R: rev}, nil
}

func decodeAddReviewResponse(_ context.Context, r *http.Response) (response interface{}, err error) {

	res := &addReviewResponse{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err
	}

	return response, nil
}
