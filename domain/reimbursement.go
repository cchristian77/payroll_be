package domain

import "time"

type Reimbursement struct {
	BaseModel

	UserID      uint64
	Description string
	Amount      uint64
	Status      string

	PayslipID    *uint64
	ReimbursedAt *time.Time

	// Associations
	Payslip *Payslip `gorm:"foreignKey:PayslipID;references:ID"`
}
