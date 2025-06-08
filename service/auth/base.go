package auth

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util/token"
	"gorm.io/gorm"
)

type Service interface {
	Authenticate(ctx context.Context, accessToken string) (*domain.User, *token.Payload, error)
	Login(ctx context.Context, input *request.Login) (*response.Auth, error)
	Logout(ctx context.Context) error
	Register(ctx context.Context) error
}

type base struct {
	repository repository.Repository
	writeDB    *gorm.DB
}

func NewService(repository repository.Repository, writerDB *gorm.DB) (Service, error) {
	return &base{
		repository: repository,
		writeDB:    writerDB,
	}, nil
}
