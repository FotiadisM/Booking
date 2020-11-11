package main

import (
	"context"

	"github.com/FotiadisM/booking/core/services/booking"
	"github.com/FotiadisM/booking/core/services/listing"
)

type repository struct {
	db []*listing.Listing
}

func newRepository() *repository {
	return &repository{}
}

func (r *repository) StoreBooking(ctx context.Context, b *booking.Booking) (err error) {

	return nil
}
