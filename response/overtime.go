package response

import (
	"github.com/cchristian77/payroll_be/domain"
	"time"
)

type Overtime struct {
	AttendanceID uint64    `json:"attendance_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Date     string `json:"date"`
	Duration uint   `json:"duration"`
}

func NewOvertimeFromDomain(o *domain.Overtime) *Overtime {
	if o == nil {
		return nil
	}

	return &Overtime{
		AttendanceID: o.AttendanceID,
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
		Date:         o.Date.Format(time.DateOnly),
		Duration:     o.Duration,
	}
}
