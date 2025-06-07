package domain

import (
	"github.com/cchristian77/payroll_be/domain/enums"
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

	// Associations
	Attendances    []Attendance    `gorm:"foreignKey:UserID"`
	Overtimes      []Overtime      `gorm:"foreignKey:UserID"`
	Reimbursements []Reimbursement `gorm:"foreignKey:UserID"`
}

func (u *User) GetHourlyRate() uint64 {
	totalHours := enums.UserWorkDays * enums.UserWorkHours

	return u.BaseSalary / totalHours
}
