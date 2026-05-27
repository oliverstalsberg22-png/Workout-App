package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string
	Email        string
	PasswordHash []byte
	Created      time.Time
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}
