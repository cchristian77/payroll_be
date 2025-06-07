package payroll_period

import (
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"time"
)

func (b *base) Upsert(ec echo.Context, input *request.UpsertPayrollPeriod) (*response.PayrollPeriod, error) {
	ctx := ec.Request().Context()
	authUser := util.EchoCntextAuthUser(ec)

	startDate, err := time.Parse(time.DateOnly, input.StartDate)
	if err != nil {
		return nil, sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("Failed on parsing start date with layout %s: %v", time.DateOnly, err))
	}

	endDate, err := time.Parse(time.DateOnly, input.EndDate)
	if err != nil {
		return nil, sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("Failed on parsing end date with layout %s: %v", time.DateOnly, err))
	}

	if startDate.After(endDate) {
		return nil, sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("the period start date %s cannot be after the end date %s",
				startDate.Format(time.DateOnly), endDate.Format(time.DateOnly)))
	}

	periods, err := b.repository.FindOverlappingPayrollPeriods(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	if len(periods) > 0 {
		return nil, sharedErrs.NewBusinessValidationErr("Payroll period is overlapping to another periods.")
	}

	// validate payroll period on update
	if input.ID != 0 {
		payrollPeriodExists, err := b.repository.FindPayrollPeriodByID(ctx, input.ID)
		if err != nil {
			return nil, err
		}

		if payrollPeriodExists == nil {
			return nil, sharedErrs.NotFoundErr
		}

		if payrollPeriodExists.PayrollRunAt != nil {
			return nil, sharedErrs.NewBusinessValidationErr(
				fmt.Sprintf("Payroll period can not be updated since the payroll is already run at %s",
					payrollPeriodExists.PayrollRunAt.Format(time.DateTime)))
		}
	}

	now := time.Now()
	payrollPeriod, err := b.repository.UpsertPayrollPeriod(ctx, &domain.PayrollPeriod{
		BaseModel: domain.BaseModel{
			ID:        input.ID,
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: authUser.ID,
			UpdatedBy: &authUser.ID,
		},
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return nil, err
	}

	return response.NewPayrollPeriodFromDomain(payrollPeriod), nil
}
