package payslip

import (
	"github.com/cchristian77/payroll_be/response"
	"github.com/labstack/echo/v4"
)

// GetSummary retrieves the summary contains the total take-home pay of all employees.
func (b *base) GetSummary(ec echo.Context, payrollPeriodID uint64) (*response.PayslipSummary, error) {
	ctx := ec.Request().Context()

	if _, err := b.EnsurePayrollExecuted(ec, payrollPeriodID); err != nil {
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
