package booking

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeConfirmPaymentEndpoint(svc ServiceModel) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(confirmPaymentRequest)
		err := svc.ConfirmPayment(ctx, req.B)
		if err != nil {
			return confirmPaymentResponse{Err: err.Error()}, nil
		}

		return confirmPaymentResponse{}, nil
	}
}
