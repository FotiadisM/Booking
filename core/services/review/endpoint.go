package review

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetAllEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		rs, err := svc.GetAll(ctx)
		if err != nil {
			return getAllResponse{nil, err.Error()}, nil
		}

		return getAllResponse{rs, ""}, nil
	}
}

func makeGetByListingIDEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByListingIDRequest)
		rs, err := svc.GetByListingID(ctx, req.ID)
		if err != nil {
			return getByListingIDResponse{nil, err.Error()}, nil
		}

		return getByListingIDResponse{rs, ""}, nil
	}
}

func makeCreateEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		id, err := svc.Create(ctx, req.R)
		if err != nil {
			return createResponse{"", err.Error()}, nil
		}

		return createResponse{id, ""}, nil
	}
}
