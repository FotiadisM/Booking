package auth

import (
	"context"
	"errors"

	"github.com/FotiadisM/booking/core/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

// ServiceModel defines the properties of a user service
type ServiceModel interface {
	AddUserLogin(context.Context, *UserLogin) (string, string, error)
	AuthUserLogin(context.Context, string, string) (string, string, error)
}

// Service containes the buisness logic
type Service struct {
	pr PersistentRepository
	tr TemporeryRepository
}

// NewService returns a Service object
func NewService() *Service {
	return &Service{}
}

var (
	// ErruserNotFound error
	ErrUserNotValidated = errors.New("Invalid user credentials")
	// ErrHashFailed error
	ErrHashFailed = errors.New("Failed to hash password")
	// ErrUserAccessTokenFailed error
	ErrUserAccessTokenFailed = errors.New("Failed to generate User Access Token")
)

func (s *Service) AddUserLogin(ctx context.Context, ul *UserLogin) (at string, rt string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(ul.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", ErrHashFailed
	}
	ul.Password = string(hash)

	if err = s.pr.AddUser(ctx, ul); err != nil {
		return
	}

	at, err = auth.CreateUserAccessToken(ul.UserID, ul.UserRole)
	if err != nil {
		return "", "", ErrUserAccessTokenFailed
	}

	return
}

func (s *Service) AuthUserLogin(ctx context.Context, Email string, Password string) (at string, rt string, err error) {
	ul, err := s.pr.GetUserByEmail(ctx, Email)
	if err != nil {
		// more error handling needed here
		return "", "", ErrUserNotValidated
	}

	if err = bcrypt.CompareHashAndPassword([]byte(ul.Password), []byte(Password)); err != nil {
		return "", "", ErrUserNotValidated
	}

	at, err = auth.CreateUserAccessToken(ul.UserID, ul.UserRole)
	if err != nil {
		return "", "", ErrUserAccessTokenFailed
	}

	return
}
