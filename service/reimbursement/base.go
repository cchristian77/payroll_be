package reimbursement

import (
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
)

type Service interface {
	Upsert(ec echo.Context, input *request.UpsertReimbursement) (*response.Reimbursement, error)
}

type base struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) (Service, error) {
	return &base{repository: repository}, nil
}
