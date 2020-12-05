package auth

import (
	"errors"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
}

// Service containes the buisness logic
type Service struct {
}

// NewService returns a Service object
func NewService() *Service {
	return &Service{}
}

var (
	// ErrNotFound returned if the user is not found
	ErrNotFound = errors.New("Listing not found")
	// ErrUUIDGen returned if uuid.NewRandom() fails
	ErrUUIDGen = errors.New("Failed to generate uuid")
)
