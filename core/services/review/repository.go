package review

import "context"

// Review defines the properties of a Listing
type Review struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	UserName  string  `json:"user_name"`
	ListingID string  `json:"listing_id"`
	Score     float32 `json:"score"`
	Comment   string  `json:"comment"`
	Created   int64   `json:"created"`
	Updated   int64   `json:"updated"`
}

// Repository describes the persistence on review model
type Repository interface {
	CreateReview(context.Context, *Review) error
	GetReviews(context.Context) ([]*Review, error)
	GetReviewsByListingID(context.Context, string) ([]*Review, error)
}
