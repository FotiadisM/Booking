package booking

import (
	"context"
	"encoding/json"
	"net/http"
)

type confirmPaymentRequest struct {
	B *Booking `json:"booking"`
}

type confirmPaymentResponse struct {
	Err string `json:"error,omitempty"`
}

func decodeConfirmPaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	b := &Booking{}
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return nil, err
	}

	return confirmPaymentRequest{B: b}, nil
}
