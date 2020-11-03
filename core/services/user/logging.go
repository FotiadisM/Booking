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

// GetUser logs info about GetUser
func (mw LoggingMiddleware) GetUser(ctx context.Context, ID string) (u *User, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	u, err = mw.Next.GetUser(ctx, ID)
	return
}

// CreateUser logs info about CreateUser
func (mw LoggingMiddleware) CreateUser(ctx context.Context, email string, firstName string, lastName string) (u *User, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	u, err = mw.Next.CreateUser(ctx, email, firstName, lastName)
	return
}
