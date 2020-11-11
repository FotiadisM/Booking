package review

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type getAllRequest struct{}

type getAllResponse struct {
	Rs  []*Review `json:"reviews,omitempty"`
	Err string    `json:"error,omitempty"`
}

// decodeGetAllRequest decodes a GetUser request
func decodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return getAllRequest{}, nil
}

type getByListingIDRequest struct {
	ID string `json:"id"`
}

type getByListingIDResponse struct {
	Rs  []*Review `json:"reviews,omitempty"`
	Err string    `json:"error,omitempty"`
}

// DecodeGetUserRequest decodes a GetUser request
func decodeGetByListingIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	return getByListingIDRequest{ID: id}, nil
}

type createRequest struct {
	R *Review
}

type createResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

// DecodeCreateUserRequest decodes a CreateUser request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	rev := &Review{}
	if err := json.NewDecoder(r.Body).Decode(rev); err != nil {
		return nil, err
	}

	return createRequest{R: rev}, nil
}

// for listing svc
type addReviewToListingRequest struct {
	ID    string  `json:"listing_id"`
	Score float32 `json:"score"`
}

type addReviewToListingResponse struct {
	Err string `json:"error,omitempty"`
}

// for serch_consumer svc
type addReviewRequest struct {
	R *Review `json:"review"`
}

type addReviewResponse struct {
	Err string `json:"error,omitempty"`
}
