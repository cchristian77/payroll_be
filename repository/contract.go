package repository

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
)

type Repository interface {

	// Session
	CreateSession(ctx context.Context, data *domain.Session) (*domain.Session, error)
	FindSessionByID(ctx context.Context, id uint64) (*domain.Session, error)
	DeleteSessionByID(ctx context.Context, id uint64) error
	RevokeSessionByID(ctx context.Context, id uint64) error

	// User
	CreateUser(ctx context.Context, data *domain.User) (*domain.User, error)
	FindUserByUsername(ctx context.Context, username string) (*domain.User, error)
	FindUserByID(ctx context.Context, id uint64) (*domain.User, error)
}
