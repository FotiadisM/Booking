package searchconsumer

import (
	"context"

	"github.com/FotiadisM/booking/core/services/listing"
	"github.com/FotiadisM/booking/core/services/review"
)

// Repository describes the persistence on user model
type Repository interface {
	AddListing(context.Context, *listing.Listing) error
	// AddBooking(context.Context, *listing.Listing) error //booking.Booking
	AddReview(context.Context, *review.Review) error
}
