package payslip

import (
	"context"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"gorm.io/gorm"
)

type Service interface {
	RunPayroll(ctx context.Context, input *request.RunPayroll) error
	FindPayslipList(ctx context.Context, input *request.FindPayslipList) (*response.BasePagination[[]*response.Payslip], error)
	FindUserPayslip(ctx context.Context, payrollPeriodID uint64) (*response.UserPayslip, error)
	GetSummary(ctx context.Context, payrollPeriodID uint64) (*response.PayslipSummary, error)
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
