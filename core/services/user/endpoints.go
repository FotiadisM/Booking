package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetByIDEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByIDRequest)
		u, err := svc.GetByID(ctx, req.ID)
		if err != nil {
			return getByIDResponse{nil, err.Error()}, nil
		}
		return getByIDResponse{u, ""}, nil
	}
}

func makeCreateEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		id, err := svc.Create(ctx, req.U)
		if err != nil {
			return createResponse{"", err.Error()}, nil
		}
		return createResponse{id, ""}, nil
	}
}
