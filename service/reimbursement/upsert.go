package reimbursement

import (
	"errors"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

func (b *base) Upsert(ec echo.Context, input *request.UpsertReimbursement) (*response.Reimbursement, error) {
	ctx := ec.Request().Context()
	authUser := util.EchoCntextAuthUser(ec)

	now := time.Now()

	// check attendance exists to create overtime
	todayAttendance, err := b.repository.FindAttendanceByUserIDAndDate(ctx, authUser.ID, now)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if todayAttendance == nil {
		return nil, sharedErrs.NewBusinessValidationErr("You must have attendance today to request the overtime.")
	}

	if todayAttendance.CheckOut == nil {
		return nil, sharedErrs.NewBusinessValidationErr("You have to finish your attendance first before requesting the overtime.")
	}

	// Check whether the reimbursement exists on update
	if input.ID != 0 {
		reimbursementExists, err := b.repository.FindReimbursementByIDAndUserID(ctx, input.ID, authUser.ID)
		if err != nil {
			return nil, err
		}

		if reimbursementExists == nil {
			return nil, sharedErrs.NotFoundErr
		}
	}

	reimbursement, err := b.repository.UpsertReimbursement(ctx, &domain.Reimbursement{
		BaseModel: domain.BaseModel{
			ID:        input.ID,
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: authUser.ID,
			UpdatedBy: &authUser.ID,
		},
		UserID:      authUser.ID,
		Description: input.Description,
		Amount:      input.Amount,
		Status:      enums.PendingReimbursementStatus,
	})
	if err != nil {
		return nil, err
	}

	return &response.Reimbursement{
		ID:          reimbursement.ID,
		CreatedAt:   reimbursement.CreatedAt,
		UpdatedAt:   reimbursement.UpdatedAt,
		Description: reimbursement.Description,
		Amount:      reimbursement.Amount,
		Status:      reimbursement.Status,
	}, nil
}
