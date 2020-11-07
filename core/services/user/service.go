package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	GetByID(context.Context, string) (*User, error)
	Create(context.Context, *User) (string, error)
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
	ErrNotFound = errors.New("User not found")
	// ErrUUIDGen returned if uuid.NewRandom() fails
	ErrUUIDGen = errors.New("Failed to generate uuid")
)

// GetByID takes the id of a user and return it from the databse
func (s *Service) GetByID(ctx context.Context, id string) (u *User, err error) {

	if u, err = s.repository.GetUserByID(ctx, id); err != nil {
		return nil, ErrNotFound
	}

	return
}

// Create creates a new user and stores it in the databse
func (s *Service) Create(ctx context.Context, u *User) (id string, err error) {

	uuid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	id = uuid.String()

	u.ID = id
	u.Created = time.Now().UTC().Unix()

	err = s.repository.CreateUser(ctx, u)

	return
}
