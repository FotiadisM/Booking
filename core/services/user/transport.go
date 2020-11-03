package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type getUserRequest struct {
	ID string `json:"id"`
}

type getUserResponse struct {
	U   *User  `json:"user,omitempty"`
	Err string `json:"error,omitempty"`
}

// DecodeGetUserRequest decodes a GetUser request
func DecodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// MakeGetUserEndpoint creates an Endpoint for GetUser
func MakeGetUserEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		u, err := svc.GetUser(ctx, req.ID)
		if err != nil {
			return getUserResponse{nil, err.Error()}, nil
		}
		return getUserResponse{u, ""}, nil
	}
}

type createUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type createUserResponse struct {
	U   *User  `json:"user,omitempty"`
	Err string `json:"error,omitempty"`
}

// DecodeCreateUserRequest decodes a CreateUser request
func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// MakeCreateUserEndpoint creates an Endpoint for CreateUser
func MakeCreateUserEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		u, err := svc.CreateUser(ctx, req.Email, req.FirstName, req.LastName)
		if err != nil {
			return createUserResponse{nil, err.Error()}, nil
		}
		return createUserResponse{u, ""}, nil
	}
}
