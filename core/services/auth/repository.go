package auth

import "context"

type User struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RepositoryPersisten describes the persistence on user model
type RepositoryPersistent interface {
	AddUser(context.Context, *User) error
	GetUser(context.Context, string) (*User, error)
}

type RepositoryTemporery interface {
	AddToken(context.Context, string) error
	GetToken(context.Context, string) (string, error)
}
