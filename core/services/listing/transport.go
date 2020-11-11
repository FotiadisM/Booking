package listing

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type getAllRequest struct{}

type getAllResponse struct {
	Ls  []*Listing `json:"listings,omitempty"`
	Err string     `json:"error,omitempty"`
}

// decodeGetAllRequest decodes a GetUser request
func decodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return getAllRequest{}, nil
}

type getByIDRequest struct {
	ID string `json:"id"`
}

type getByIDResponse struct {
	L   *Listing `json:"listing,omitempty"`
	Err string   `json:"error,omitempty"`
}

// DecodeGetUserRequest decodes a GetUser request
func decodeGetByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	return getByIDRequest{ID: id}, nil
}

type createRequest struct {
	L *Listing
}

type createResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

// DecodeCreateUserRequest decodes a CreateUser request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	l := &Listing{}
	if err := json.NewDecoder(r.Body).Decode(l); err != nil {
		return nil, err
	}

	return createRequest{L: l}, nil
}

type addReviewToListingRequest struct {
	ID    string  `json:"listing_id"`
	Score float32 `json:"score"`
}

type addReviewToListingResponse struct {
	Err string `json:"error,omitempty"`
}

// DecodeCreateUserRequest decodes a CreateUser request
func decodeAddReviewToListingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var m map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		return nil, err
	}

	return addReviewToListingRequest{ID: m["listing_id"].(string), Score: m["score"].(float32)}, nil
}

// DecodeCreateUserResponse decodes a CreateUser request
func decodeAddReviewToListingRespnose(_ context.Context, r *http.Response) (interface{}, error) {
	var res addReviewToListingResponse

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

type addListingRequest struct {
	L *Listing `json:"listing"`
}

type addListingResponse struct {
	Err string `json:"error,omitempty"`
}
