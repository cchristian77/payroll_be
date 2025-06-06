package attendance

import (
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
)

type Service interface {
	CheckIn(ec echo.Context) (*response.Attendance, error)
	CheckOut(ec echo.Context) (*response.Attendance, error)
}

type base struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) (Service, error) {
	return &base{repository: repository}, nil
}
