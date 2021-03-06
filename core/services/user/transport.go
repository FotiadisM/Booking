package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type getByIDRequest struct {
	ID string `json:"id"`
}

type getByIDResponse struct {
	U   *User  `json:"user,omitempty"`
	Err string `json:"error,omitempty"`
}

// DecodeGetUserRequest decodes a GetUser request
func decodeGetByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	req := getByIDRequest{ID: id}

	return req, nil
}

type createRequest struct {
	U *User
}

type createResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

// DecodeCreateUserRequest decodes a CreateUser request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	u := &User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return nil, err
	}

	return createRequest{U: u}, nil
}
