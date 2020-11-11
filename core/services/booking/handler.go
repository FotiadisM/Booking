package booking

import "github.com/go-kit/kit/transport/http"

// ConfirmPaymentHandler returns a http Handler
func ConfirmPaymentHandler(svc ServiceModel) *http.Server {
	return http.NewServer(
		makeConfirmPaymentEndpoint(svc),
		decodeConfirmPaymentRequest,
		http.EncodeJSONResponse,
	)
}
