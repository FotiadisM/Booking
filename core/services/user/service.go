package user

import (
	"context"
	"errors"
	"time"
)

// User defines the properties of a User
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Created   int64  `json:"-"`
}

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	GetUser(context.Context, string) (*User, error)
	CreateUser(context.Context, string, string, string) (*User, error)
}

// Service containes the buisness logic
type Service struct{}

var (
	// ErrNotFound returned if the user is not found
	ErrNotFound = errors.New("User not found")
)

var userRepo []*User

// GetUser takes the id of a user and return it from the databse
func (Service) GetUser(ctx context.Context, ID string) (*User, error) {
	for _, v := range userRepo {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, ErrNotFound
}

// CreateUser creates a new user and stores it in the databse
func (Service) CreateUser(ctx context.Context, email string, firstName string, lastName string) (*User, error) {
	u := &User{
		ID:        "1234",
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Created:   time.Now().UTC().Unix(),
	}
	userRepo = append(userRepo, u)

	return u, nil
}
