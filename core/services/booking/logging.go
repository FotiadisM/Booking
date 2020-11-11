package booking

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

// ConfirmPayment logs info about GetUser
func (mw LoggingMiddleware) ConfirmPayment(ctx context.Context, b *Booking) (err error) {

	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.ConfirmPayment(ctx, b)
	return
}
