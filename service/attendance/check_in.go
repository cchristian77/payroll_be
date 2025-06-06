package attendance

import (
	"errors"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

func (b base) CheckIn(ec echo.Context) (*response.Attendance, error) {
	ctx := ec.Request().Context()
	authUser := util.EchoCntextAuthUser(ec)

	now := time.Now()

	attendanceExists, err := b.repository.FindAttendanceByUserIDAndDate(ctx, authUser.ID, now)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if attendanceExists != nil {
		return nil, sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("You have already checked in at %s", attendanceExists.CheckIn.Format("2006-01-02 15:04:05")),
		)
	}

	attendance, err := b.repository.CreateAttendance(ctx, &domain.Attendance{
		BaseModel: domain.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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

	return &response.Attendance{
		AttendanceID: attendance.ID,
		CreatedAt:    attendance.CreatedAt,
		UpdatedAt:    attendance.UpdatedAt,
		Date:         attendance.Date,
		CheckIn:      attendance.CheckIn,
		CheckOut:     attendance.CheckOut,
	}, nil
}
