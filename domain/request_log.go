package domain

import "time"

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
