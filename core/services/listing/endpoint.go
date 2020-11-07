package listing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetAllEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ls, err := svc.GetAll(ctx)
		if err != nil {
			return getAllResponse{nil, err.Error()}, nil
		}
		return getAllResponse{ls, ""}, nil
	}
}

func makeGetByIDEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByIDRequest)
		l, err := svc.GetByID(ctx, req.ID)
		if err != nil {
			return getByIDResponse{nil, err.Error()}, nil
		}
		return getByIDResponse{l, ""}, nil
	}
}

func makeCreateEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		id, err := svc.Create(ctx, req.L)
		if err != nil {
			return createResponse{"", err.Error()}, nil
		}
		return createResponse{id, ""}, nil
	}
}
