package listing

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	GetAll(context.Context) ([]*Listing, error)
	GetByID(context.Context, string) (*Listing, error)
	Create(context.Context, *Listing) (string, error)
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

// GetAll returns all listings
func (s *Service) GetAll(ctx context.Context) (listings []*Listing, err error) {
	if listings, err = s.repository.GetListings(ctx); err != nil {
		return nil, err
	}

	return
}

// GetByID returns the listing with ID == id
func (s *Service) GetByID(ctx context.Context, id string) (l *Listing, err error) {
	if l, err = s.repository.GetListingByID(ctx, id); err != nil {
		return nil, ErrNotFound
	}

	return
}

// Create creates a new listing and adds it to the databse
func (s *Service) Create(ctx context.Context, l *Listing) (id string, err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	id = uuid.String()

	l.ID = id
	l.Created = time.Now().UTC().Unix()

	if err = s.repository.CreateListing(ctx, l); err != nil {
		return "", err
	}

	return
}
