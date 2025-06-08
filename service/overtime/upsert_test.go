package overtime

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/cchristian77/payroll_be/util/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func (suite *OvertimeServiceTestSuite) Test_Upsert() {
	attendance := mock.InitAttendanceDomain()
	overtime := mock.InitOvertimeDomain()
	input := &request.UpsertOvertime{Duration: 1}

	var expected *response.Overtime

	testCases := []struct {
		name          string
		prepareMock   func()
		expectedError error
		wantErr       bool
	}{
		{
			name: "attendance not found",
			prepareMock: func() {
				authUser := util.AuthUserFromCtx(suite.ctx)

				suite.repo.EXPECT().FindAttendanceByUserIDAndDate(suite.ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(nil, gorm.ErrRecordNotFound).
					Times(1)
				suite.repo.EXPECT().UpsertOvertime(suite.ctx, gomock.Any()).Times(0)
			},
			wantErr:       true,
			expectedError: sharedErrs.NewBusinessValidationErr("You must have attendance today to request the overtime."),
		},
		{
			name: "not checked out yet",
			prepareMock: func() {
				authUser := util.AuthUserFromCtx(suite.ctx)

				attendance.CheckOut = nil

				suite.repo.EXPECT().FindAttendanceByUserIDAndDate(suite.ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(attendance, nil).
					Times(1)
				suite.repo.EXPECT().UpsertOvertime(suite.ctx, gomock.Any()).Times(0)
			},
			wantErr:       true,
			expectedError: sharedErrs.NewBusinessValidationErr("You have to finish your attendance first before requesting the overtime."),
		},
		{
			name: "success",
			prepareMock: func() {
				authUser := util.AuthUserFromCtx(suite.ctx)

				attendance.CheckOut = &attendance.CheckIn
				expected = response.NewOvertimeFromDomain(overtime)

				suite.repo.EXPECT().FindAttendanceByUserIDAndDate(suite.ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(attendance, nil).
					Times(1)
				suite.repo.EXPECT().UpsertOvertime(suite.ctx, gomock.Any()).
					Return(overtime, nil).
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
			result, err := suite.overtimeService.Upsert(suite.ctx, input)

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
