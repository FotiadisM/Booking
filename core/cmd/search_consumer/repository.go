package main

import (
	"context"

	"github.com/FotiadisM/booking/core/services/listing"
	"github.com/FotiadisM/booking/core/services/review"
)

type repository struct {
}

func newRepository() *repository {
	return &repository{}
}

func (r *repository) AddListing(ctx context.Context, l *listing.Listing) (err error) {

	return
}

func (r *repository) AddReview(ctx context.Context, rev *review.Review) (err error) {

	return
}
