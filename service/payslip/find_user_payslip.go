package payslip

import (
	"context"
	"errors"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"gorm.io/gorm"
)

func (b *base) FindUserPayslip(ctx context.Context, payrollPeriodID uint64) (*response.UserPayslip, error) {
	authUser := util.AuthUserFromCtx(ctx)

	payrollPeriod, err := b.EnsurePayrollExecuted(ctx, payrollPeriodID)
	if err != nil {
		return nil, err
	}

	payslip, err := b.repository.FindPayslipByUserIDAndPayrollPeriodID(ctx, authUser.ID, payrollPeriodID)
	if err != nil {
		return nil, err
	}
	payslip.PayrollPeriod = payrollPeriod

	reimbursements, err := b.repository.FindReimbursementsByPayslipID(ctx, payslip.ID)
	if err != nil {
		return nil, err
	}
	payslip.Reimbursements = reimbursements

	return response.NewUserPayslipFromDomain(payslip), nil
}

func (b *base) EnsurePayrollExecuted(ctx context.Context, payrollPeriodID uint64) (*domain.PayrollPeriod, error) {
	payrollPeriod, err := b.repository.FindPayrollPeriodByID(ctx, payrollPeriodID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if payrollPeriod == nil {
		return nil, sharedErrs.NewBusinessValidationErr("Payroll period not found.")
	}

	if payrollPeriod.PayrollRunAt == nil {
		return nil, sharedErrs.NewBusinessValidationErr("Payroll period has not executed yet.")
	}

	return payrollPeriod, nil
}
