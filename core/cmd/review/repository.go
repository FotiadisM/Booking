package main

import (
	"context"
	"errors"

	"github.com/FotiadisM/booking/core/services/review"
)

type repository struct {
	db []*review.Review
}

func newRepository() *repository {
	return &repository{}
}

func (r *repository) GetReviews(ctx context.Context) (rs []*review.Review, err error) {

	return r.db, nil
}

func (r *repository) CreateReview(ctx context.Context, rev *review.Review) (err error) {
	r.db = append(r.db, rev)

	return
}

func (r *repository) GetReviewsByListingID(ctx context.Context, id string) (rs []*review.Review, err error) {
	for _, v := range r.db {
		if v.ListingID == id {
			rs = append(rs, v)
		}
	}

	if len(rs) == 0 {
		return nil, errors.New("Not found")
	}

	return
}
