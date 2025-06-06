package domain

import "time"

type BaseModel struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uint64
	UpdatedBy *uint64
}
