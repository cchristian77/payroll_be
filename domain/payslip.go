package domain

type Payslip struct {
	BaseModel

	PayslipPeriodID    uint64
	UserID             uint64
	TotalAttendance    uint
	TotalOvertime      uint
	TotalReimbursement uint
	BaseSalary         uint64
	TotalSalary        uint64
}
