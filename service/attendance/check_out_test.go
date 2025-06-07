package attendance

import (
	"fmt"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/cchristian77/payroll_be/util/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
	"time"
)

func (suite *AttendanceServiceTestSuite) Test_CheckOut() {
	attendance := mock.InitAttendanceDomain()

	var expected *response.Attendance

	testCases := []struct {
		name          string
		prepareMock   func()
		expectedError error
		wantErr       bool
	}{
		{
			name: "attendance not found",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()
				authUser := util.EchoCntextAuthUser(suite.ec)

				suite.repo.EXPECT().FindAttendanceByUserIDAndDate(ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(nil, gorm.ErrRecordNotFound).
					Times(1)
				suite.repo.EXPECT().UpdateAttendance(ctx, gomock.Any()).Times(0)
			},
			wantErr:       true,
			expectedError: sharedErrs.NewBusinessValidationErr(fmt.Sprintf("You haven't checked in yet today.")),
		},
		{
			name: "already checked out",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()
				authUser := util.EchoCntextAuthUser(suite.ec)

				attendance.CheckOut = &attendance.CheckIn

				suite.repo.EXPECT().FindAttendanceByUserIDAndDate(ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(attendance, nil).
					Times(1)
				suite.repo.EXPECT().UpdateAttendance(ctx, gomock.Any()).Times(0)
			},
			wantErr: true,
			expectedError: sharedErrs.NewBusinessValidationErr(
				fmt.Sprintf("You have already checked out at %s", attendance.CheckIn.Format(time.DateTime))),
		},
		{
			name: "success",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()
				authUser := util.EchoCntextAuthUser(suite.ec)

				attendance.CheckOut = nil
				expected = response.NewAttendanceFromDomain(attendance)

				suite.repo.EXPECT().FindAttendanceByUserIDAndDate(ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(attendance, nil).
					Times(1)
				suite.repo.EXPECT().UpdateAttendance(ctx, gomock.Any()).
					Return(nil).
					Times(1)
				suite.repo.EXPECT().FindAttendanceByIDAndUserID(ctx, gomock.Eq(attendance.ID), gomock.Eq(authUser.ID)).
					Return(attendance, nil).
					Times(1)
			},
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			// Arrange
			suite.Before(t)
			defer suite.After(t)
			tc.prepareMock()

			// Act
			result, err := suite.attendanceService.CheckOut(suite.ec)

			// Assert
			assert.Equal(t, tc.wantErr, err != nil, "error expected %v, but actual: %v", tc.wantErr, err)
			if tc.wantErr {
				assert.Empty(t, result)
				assert.Equal(t, tc.expectedError, err, "error expected %v, but actual: %v", tc.expectedError, err)
			} else {
				assert.NotEmpty(t, result)
				if err = util.CompareData(result, expected, 1); err != nil {
					t.Fatalf("error on comparing data : %v", err)
				}
			}
		})
	}
}
