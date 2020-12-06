package auth

import "context"

type UserLogin struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RepositoryPersisten describes the persistence on user model
type PersistentRepository interface {
	AddUser(context.Context, *UserLogin) error
	GetUserByEmail(context.Context, string) (*UserLogin, error)
}

type TemporeryRepository interface {
	AddToken(context.Context, string) error
	GetToken(context.Context, string) (string, error)
}
