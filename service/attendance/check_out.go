package attendance

import (
	"context"
	"errors"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"gorm.io/gorm"
	"time"
)

func (b *base) CheckOut(ctx context.Context) (*response.Attendance, error) {
	authUser := util.AuthUserFromCtx(ctx)

	now := time.Now()

	attendance, err := b.repository.FindAttendanceByUserIDAndDate(ctx, authUser.ID, now)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// before checking out, the user should checkin first.
	if attendance == nil {
		return nil, sharedErrs.NewBusinessValidationErr(fmt.Sprintf("You haven't checked in yet today."))
	}

	// validated if user's already checkout
	if attendance.CheckOut != nil {
		return nil, sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("You have already checked out at %s", attendance.CheckOut.Format(time.DateTime)),
		)
	}

	err = b.repository.UpdateAttendance(ctx, &domain.Attendance{
		BaseModel: domain.BaseModel{
			ID:        attendance.ID,
			UpdatedAt: now,
			UpdatedBy: &authUser.ID,
		},
		CheckOut: &now,
	})
	if err != nil {
		return nil, err
	}

	attendance, err = b.repository.FindAttendanceByIDAndUserID(ctx, attendance.ID, authUser.ID)
	if err != nil {
		return nil, err
	}

	return response.NewAttendanceFromDomain(attendance), nil
}
