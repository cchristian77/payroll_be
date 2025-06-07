package payroll_period

import (
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
	Upsert(ec echo.Context, input *request.UpsertPayrollPeriod) (*response.PayrollPeriod, error)
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
