package main

import (
	"context"
	"errors"

	"github.com/FotiadisM/booking/core/services/user"
)

type repository struct {
	db []*user.User
}

func newRepository() *repository {
	return &repository{}
}

func (r *repository) CreateUser(ctx context.Context, u *user.User) (err error) {
	r.db = append(r.db, u)

	return
}

func (r *repository) GetUserByID(ctx context.Context, id string) (u *user.User, err error) {
	for _, v := range r.db {
		if v.ID == id {
			u = v
			return
		}
	}

	return nil, errors.New("Not found")
}
