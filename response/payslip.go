package response

import (
	"github.com/cchristian77/payroll_be/domain"
	"time"
)

type Payslip struct {
	ID uint64 `json:"id"`

	PayrollPeriodID     uint64 `json:"payroll_period_id"`
	PayrollStartDate    string `json:"payroll_start_date"`
	PayrollEndDate      string `json:"payroll_end_date"`
	UserID              uint64 `json:"user_id"`
	FullName            string `json:"full_name"`
	TotalAttendanceDays uint   `json:"total_attendance_days"`
	TotalOvertimeDays   uint   `json:"total_overtime_days"`
	TotalOvertimeHours  uint   `json:"total_overtime_hours"`
	TotalReimbursements uint64 `json:"total_reimbursements"`
	BaseSalary          uint64 `json:"base_salary"`
	AttendancePay       uint64 `json:"attendance_pay"`
	OvertimePay         uint64 `json:"overtime_pay"`
	ReimbursementPay    uint64 `json:"reimbursement_pay"`
	TakeHomePay         uint64 `json:"take_home_pay"`
}

type PayslipSummary struct {
	TotalTakeHomePay uint64 `json:"total_take_home_pay"`
}

func NewPayslipFromDomain(p *domain.Payslip) *Payslip {
	if p == nil {
		return nil
	}

	result := &Payslip{
		ID:                  p.ID,
		UserID:              p.UserID,
		TotalAttendanceDays: p.TotalAttendanceDays,
		TotalOvertimeDays:   p.TotalOvertimeDays,
		TotalOvertimeHours:  p.TotalOvertimeHours,
		TotalReimbursements: p.TotalReimbursements,
		BaseSalary:          p.BaseSalary,
		AttendancePay:       p.AttendancePay,
		OvertimePay:         p.OvertimePay,
		ReimbursementPay:    p.ReimbursementPay,
		TakeHomePay:         p.TotalSalary,
	}

	if p.PayrollPeriod != nil {
		result.PayrollPeriodID = p.PayrollPeriod.ID
		result.PayrollStartDate = p.PayrollPeriod.StartDate.Format(time.DateOnly)
		result.PayrollEndDate = p.PayrollPeriod.EndDate.Format(time.DateOnly)
	}

	if p.User != nil {
		result.FullName = p.User.FullName
	}

	return result
}
