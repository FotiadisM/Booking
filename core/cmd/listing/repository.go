package main

import (
	"context"
	"errors"

	"github.com/FotiadisM/booking/core/services/listing"
)

type repository struct {
	db []*listing.Listing
}

func newRepository() *repository {
	return &repository{}
}

func (r *repository) GetListings(ctx context.Context) (listings []*listing.Listing, err error) {

	return r.db, nil
}

func (r *repository) CreateListing(ctx context.Context, l *listing.Listing) (err error) {
	r.db = append(r.db, l)

	return
}

func (r *repository) GetListingByID(ctx context.Context, id string) (l *listing.Listing, err error) {
	for _, v := range r.db {
		if v.ID == id {
			l = v
			return
		}
	}

	return nil, errors.New("Listing not found")
}

func (r *repository) UpdateListing(ctx context.Context, l *listing.Listing) (err error) {
	for i := range r.db {
		if r.db[i].ID == l.ID {
			r.db[i] = l
			return
		}
	}

	return errors.New("Listing not found")
}
