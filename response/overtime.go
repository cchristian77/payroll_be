package response

import "time"

type Overtime struct {
	AttendanceID uint64    `json:"attendance_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Date     string `json:"date"`
	Duration uint   `json:"duration"`
}
