package response

import (
	"github.com/cchristian77/payroll_be/domain"
	"time"
)

type Attendance struct {
	AttendanceID uint64    `json:"attendance_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Date     string     `json:"date"`
	CheckIn  time.Time  `json:"check_in"`
	CheckOut *time.Time `json:"check_out"`
}

func NewAttendanceFromDomain(a *domain.Attendance) *Attendance {
	if a == nil {
		return nil
	}

	return &Attendance{
		AttendanceID: a.ID,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
		Date:         a.Date.Format(time.DateOnly),
		CheckIn:      a.CheckIn,
		CheckOut:     a.CheckOut,
	}
}
