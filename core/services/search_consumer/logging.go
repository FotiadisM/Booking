package searchconsumer

import (
	"context"
	"time"

	"github.com/FotiadisM/booking/core/services/listing"
	"github.com/FotiadisM/booking/core/services/review"
	"github.com/go-kit/kit/log"
)

// LoggingMiddleware logs stuff
type LoggingMiddleware struct {
	Logger log.Logger
	Next   ServiceModel
}

// AddListing logs info about GetUser
func (mw LoggingMiddleware) AddListing(ctx context.Context, l *listing.Listing) (err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.AddListing(ctx, l)
	return
}

// AddReview logs info about GetUser
func (mw LoggingMiddleware) AddReview(ctx context.Context, r *review.Review) (err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.AddReview(ctx, r)
	return
}
