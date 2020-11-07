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

	req := getByIDRequest{ID: id}

	return req, nil
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
