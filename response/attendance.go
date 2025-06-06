package response

import "time"

type Attendance struct {
	AttendanceID uint64    `json:"attendance_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Date     time.Time  `json:"date"`
	CheckIn  time.Time  `json:"check_in"`
	CheckOut *time.Time `json:"check_out"`
}
