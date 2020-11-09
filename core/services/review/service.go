package review

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	GetAll(context.Context) ([]*Review, error)
	GetByListingID(context.Context, string) ([]*Review, error)
	Create(context.Context, *Review) (string, error)
}

// Service containes the buisness logic
type Service struct {
	repository Repository
}

// NewService returns a Service object
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

var (
	// ErrNotFound returned if the user is not found
	ErrNotFound = errors.New("Listing not found")
	// ErrUUIDGen returned if uuid.NewRandom() fails
	ErrUUIDGen = errors.New("Failed to generate uuid")
)

// GetAll returns all reviews
func (s *Service) GetAll(ctx context.Context) (rs []*Review, err error) {
	if rs, err = s.repository.GetReviews(ctx); err != nil {
		return nil, err
	}

	return
}

// GetByListingID returns all reviews with review.listingID == id
func (s *Service) GetByListingID(ctx context.Context, id string) (rs []*Review, err error) {
	if rs, err = s.repository.GetReviewsByListingID(ctx, id); err != nil {
		return nil, ErrNotFound
	}

	return
}

// Create creates a new review and adds it to the databse
func (s *Service) Create(ctx context.Context, r *Review) (id string, err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	id = uuid.String()

	r.ID = id
	r.Created = time.Now().UTC().Unix()

	if err = s.repository.CreateReview(ctx, r); err != nil {
		return "", err
	}

	return
}
