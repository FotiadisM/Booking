package user

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

// GetByID logs info about GetUser
func (mw LoggingMiddleware) GetByID(ctx context.Context, id string) (u *User, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	u, err = mw.Next.GetByID(ctx, id)
	return
}

// Create logs info about CreateUser
func (mw LoggingMiddleware) Create(ctx context.Context, u *User) (id string, err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	id, err = mw.Next.Create(ctx, u)

	return
}
