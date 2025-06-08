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

func (b *base) CheckIn(ctx context.Context) (*response.Attendance, error) {
	authUser := util.AuthUserFromCtx(ctx)

	now := time.Now()

	// user can not check in during the weekend
	if util.Contains([]time.Weekday{time.Saturday, time.Sunday}, now.Weekday()) {
		return nil, sharedErrs.NewBusinessValidationErr("You are not able to check on the weekend.")
	}

	attendanceExists, err := b.repository.FindAttendanceByUserIDAndDate(ctx, authUser.ID, now)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// if attendance exists for today, it implies that user's already checked in today.
	if attendanceExists != nil {
		return nil, sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("You have already checked in %s", attendanceExists.CheckIn.Format(time.DateTime)),
		)
	}

	attendance, err := b.repository.CreateAttendance(ctx, &domain.Attendance{
		BaseModel: domain.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: authUser.ID,
		},
		UserID:   authUser.ID,
		Date:     now,
		CheckIn:  now,
		CheckOut: nil,
	})
	if err != nil {
		return nil, err
	}

	return response.NewAttendanceFromDomain(attendance), nil
}
