package payroll_period

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func (suite *PayrollPeriodTestSuite) Test_Upsert() {
	now := time.Now()
	input := &request.UpsertPayrollPeriod{
		ID:        1,
		StartDate: now.Format(time.DateOnly),
		EndDate:   now.Add(30 * 24 * time.Hour).Format(time.DateOnly),
	}

	var expected *response.PayrollPeriod

	testCases := []struct {
		name          string
		prepareMock   func()
		expectedError error
		wantErr       bool
	}{
		{
			name: "invalid start date",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()

				input.StartDate = "invalid_date"

				suite.repo.EXPECT().UpsertPayrollPeriod(ctx, gomock.Any()).Times(0)
			},
			wantErr: true,
		},
		{
			name: "invalid end date date",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()

				input.StartDate = now.Format(time.DateOnly)
				input.EndDate = "invalid_date"

				suite.repo.EXPECT().UpsertPayrollPeriod(ctx, gomock.Any()).Times(0)
			},
			wantErr: true,
		},
		{
			name: "start date is after end date",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()

				input.StartDate = now.Format(time.DateOnly)
				input.EndDate = now.Add(-30 * 24 * time.Hour).Format(time.DateOnly)

				suite.repo.EXPECT().UpsertPayrollPeriod(ctx, gomock.Any()).Times(0)
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			// Arrange
			suite.Before(t)
			defer suite.After(t)
			tc.prepareMock()

			// Act
			result, err := suite.payrollPeriodService.Upsert(suite.ec, input)

			// Assert
			assert.Equal(t, tc.wantErr, err != nil, "error expected %v, but actual: %v", tc.wantErr, err)
			if tc.wantErr {
				assert.Empty(t, result)
				assert.Error(t, err)
			} else {
				assert.NotEmpty(t, result)
				if err = util.CompareData(result, expected, 1); err != nil {
					t.Fatalf("error on comparing data : %v", err)
				}
			}
		})
	}
}
