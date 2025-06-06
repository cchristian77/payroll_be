package auth

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
)

type Service interface {
	Authenticate(ctx context.Context, token string) (*domain.User, error)
	Login(ctx context.Context, input *request.Login) (*response.Auth, error)
	Logout(ctx context.Context, sessionID uint64) error
}

type base struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) (Service, error) {
	return &base{repository: repository}, nil
}
