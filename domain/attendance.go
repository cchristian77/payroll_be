package domain

import "time"

type Attendance struct {
	BaseModel

	UserID   uint64
	Date     time.Time
	CheckIn  time.Time
	CheckOut *time.Time

	// Associations
	Overtime *Overtime `gorm:"foreignKey:AttendanceID"`
}
