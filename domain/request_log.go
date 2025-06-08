package domain

import "time"

// RequestLog represents a log entry for tracking user or system activities performed in the application.
type RequestLog struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time

	RequestID string
	UserID    uint64
	Activity  string
	Entity    string
	RefID     uint64
	ClientIP  string
}
