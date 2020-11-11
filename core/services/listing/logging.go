package listing

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware logs stuff
type LoggingMiddleware struct {
	Logger log.Logger
	Next   ServiceModel
}

// GetAll logs info about GetUser
func (mw LoggingMiddleware) GetAll(ctx context.Context) (listings []*Listing, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	listings, err = mw.Next.GetAll(ctx)
	return
}

// GetByID logs info about GetUser
func (mw LoggingMiddleware) GetByID(ctx context.Context, id string) (l *Listing, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	l, err = mw.Next.GetByID(ctx, id)
	return
}

// Create logs info about CreateUser
func (mw LoggingMiddleware) Create(ctx context.Context, l *Listing) (id string, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	id, err = mw.Next.Create(ctx, l)

	return
}

// AddReviewToListing logs info about CreateUser
func (mw LoggingMiddleware) AddReviewToListing(ctx context.Context, id string, score float32) (err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.AddReviewToListing(ctx, id, score)

	return
}
