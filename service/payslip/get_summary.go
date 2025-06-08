package payslip

import (
	"context"
	"github.com/cchristian77/payroll_be/response"
)

// GetSummary retrieves the summary contains the total take-home pay of all employees.
func (b *base) GetSummary(ctx context.Context, payrollPeriodID uint64) (*response.PayslipSummary, error) {
	if _, err := b.EnsurePayrollExecuted(ctx, payrollPeriodID); err != nil {
		return nil, err
	}

	totalTakeHomePay, err := b.repository.FindPayslipSumTotalSalary(ctx, payrollPeriodID)
	if err != nil {
		return nil, err
	}

	return &response.PayslipSummary{
		TotalTakeHomePay: totalTakeHomePay,
	}, nil
}
