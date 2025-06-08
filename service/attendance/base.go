package attendance

import (
	"context"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/response"
	"gorm.io/gorm"
)

type Service interface {
	CheckIn(ctx context.Context) (*response.Attendance, error)
	CheckOut(ctx context.Context) (*response.Attendance, error)
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
