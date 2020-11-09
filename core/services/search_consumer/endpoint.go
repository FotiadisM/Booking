package searchconsumer

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeAddListingEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addListingRequest)
		err := svc.AddListing(ctx, req.L)
		if err != nil {
			return addListingResponse{err.Error()}, nil
		}

		return addListingResponse{""}, nil
	}
}

func makeAddReviewEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addReviewRequest)
		err := svc.AddReview(ctx, req.R)
		if err != nil {
			return addReviewResponse{err.Error()}, nil
		}

		return addReviewResponse{""}, nil
	}
}
