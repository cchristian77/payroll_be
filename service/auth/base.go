package auth

import (
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
	Authenticate(ec echo.Context, token string) (*domain.User, error)
	Login(ec echo.Context, input *request.Login) (*response.Auth, error)
	Logout(ec echo.Context, sessionID uint64) error
	Register(ec echo.Context) error
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
