package payslip

import (
	"errors"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (b *base) GetSummary(ec echo.Context, input *request.RunPayroll) (*response.PayslipSummary, error) {
	ctx := ec.Request().Context()

	payrollPeriod, err := b.repository.FindPayrollPeriodByID(ctx, input.PayrollPeriodID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if payrollPeriod == nil {
		return nil, sharedErrs.NewBusinessValidationErr("Payroll period not found.")
	}

	totalTakeHomePay, err := b.repository.FindPayslipSumTotalSalary(ctx, payrollPeriod.ID)
	if err != nil {
		return nil, err
	}

	return &response.PayslipSummary{
		TotalTakeHomePay: totalTakeHomePay,
	}, nil
}
