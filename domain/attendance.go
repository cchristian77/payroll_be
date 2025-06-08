package domain

import "time"

// Attendance represents the attendance record of a user for a specific date.
type Attendance struct {
	BaseModel

	UserID   uint64
	Date     time.Time
	CheckIn  time.Time
	CheckOut *time.Time

	// Associations
	Overtime *Overtime `gorm:"foreignKey:AttendanceID"`
}
