package review

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
func (mw LoggingMiddleware) GetAll(ctx context.Context) (rs []*Review, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	rs, err = mw.Next.GetAll(ctx)
	return
}

// GetByListingID logs info about GetUser
func (mw LoggingMiddleware) GetByListingID(ctx context.Context, id string) (rs []*Review, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	rs, err = mw.Next.GetByListingID(ctx, id)
	return
}

// Create logs info about CreateUser
func (mw LoggingMiddleware) Create(ctx context.Context, r *Review) (id string, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	id, err = mw.Next.Create(ctx, r)

	return
}
