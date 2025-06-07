package domain

type Payslip struct {
	BaseModel

	PayrollPeriodID     uint64
	UserID              uint64
	TotalAttendanceDays uint
	TotalOvertimeDays   uint
	TotalOvertimeHours  uint
	TotalReimbursements uint64
	BaseSalary          uint64
	AttendancePay       uint64
	OvertimePay         uint64
	ReimbursementPay    uint64
	TotalSalary         uint64

	// Associations
	User           *User            `gorm:"foreignKey:UserID;references:ID"`
	PayrollPeriod  *PayrollPeriod   `gorm:"foreignKey:PayrollPeriodID;references:ID"`
	Reimbursements []*Reimbursement `gorm:"foreignKey:PayslipID;references:ID"`
}
