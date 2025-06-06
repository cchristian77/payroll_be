package overtime

import (
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
)

type Service interface {
	Upsert(ec echo.Context, input *request.UpsertOvertime) (*response.Overtime, error)
}

type base struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) (Service, error) {
	return &base{repository: repository}, nil
}
