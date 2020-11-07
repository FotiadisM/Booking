package user

import "context"

// Role is an enum for the user role
type Role int

const (
	// Admin h
	Admin Role = iota

	// Host is hosting apartments on the website and can also rent
	Host

	// Tenant is renting apartments
	Tenant
)

// User defines the properties of a User
type User struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Role            Role   `json:"role"`
	TelephoneNumber string `json:"tel_number"`
	Created         int64  `json:"-"`
}

// Repository describes the persistence on user model
type Repository interface {
	CreateUser(context.Context, *User) error
	GetUserByID(context.Context, string) (*User, error)
}
