package request_log

import (
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
	Log(ec echo.Context, activity string, referenceID uint64, entity string) (*domain.RequestLog, error)
}

type base struct {
	repository repository.Repository
	writeDB    *gorm.DB
}

func NewService(repository repository.Repository, writeDB *gorm.DB) (Service, error) {
	return &base{
		repository: repository,
		writeDB:    writeDB,
	}, nil
}
