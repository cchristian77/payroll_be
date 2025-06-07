package payslip

import (
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
	RunPayroll(ec echo.Context, input *request.RunPayroll) error
	FindPayslipList(ec echo.Context, input *request.FindPayslipList) (*response.BasePagination[[]*response.Payslip], error)
	GetSummary(ec echo.Context, input *request.RunPayroll) (*response.PayslipSummary, error)
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
