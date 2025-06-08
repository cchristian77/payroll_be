package overtime

import (
	"context"
	"errors"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"gorm.io/gorm"
	"time"
)

func (b *base) Upsert(ctx context.Context, input *request.UpsertOvertime) (*response.Overtime, error) {
	authUser := util.AuthUserFromCtx(ctx)

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

	overtime, err := b.repository.UpsertOvertime(ctx, &domain.Overtime{
		AttendanceID: todayAttendance.ID,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    authUser.ID,
		UpdatedBy:    &authUser.ID,

		UserID:   authUser.ID,
		Date:     now,
		Duration: input.Duration,
	})
	if err != nil {
		return nil, err
	}

	return response.NewOvertimeFromDomain(overtime), nil
}
