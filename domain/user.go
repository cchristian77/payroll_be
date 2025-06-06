package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt

	Username   string
	FullName   string
	Password   string
	Role       string
	BaseSalary uint64
}
