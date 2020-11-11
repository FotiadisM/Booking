package searchconsumer

import (
	"context"
	"fmt"

	"github.com/FotiadisM/booking/core/services/listing"
	"github.com/FotiadisM/booking/core/services/review"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	AddListing(context.Context, *listing.Listing) error
	AddReview(context.Context, *review.Review) error
	// AddBooking(context.Context, booking.Booking) error
}

// Service is an implementation of the ServiceModel
type Service struct {
	repository Repository
}

// NewService returns a Service object
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// AddListing adds the newly created listing to the search databse
func (s *Service) AddListing(context.Context, *listing.Listing) (err error) {
	fmt.Println("ADD LISTING")

	return
}

// AddReview adds the newly created review to the search databse
func (s *Service) AddReview(context.Context, *review.Review) (err error) {
	fmt.Println("ADD REVIEW")

	return
}
