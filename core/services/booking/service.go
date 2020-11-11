package booking

import (
	"context"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	ConfirmPayment(context.Context, *Booking) error
}

// Service containes the buisness logic
type Service struct {
	repository Repository
}

// NewService returns a Service object
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// ConfirmPayment rfs
func (s *Service) ConfirmPayment(ctx context.Context, b *Booking) (err error) {

	return
}
