package payslip

import (
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	"github.com/cchristian77/payroll_be/util/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func (suite *PayslipTestSuite) Test_GetSummary() {
	payrollPeriodID := uint64(1)
	payrollPeriod := mock.InitPayrollPeriodDomain()

	var expected *response.PayslipSummary

	testCases := []struct {
		name          string
		prepareMock   func()
		expectedError error
		wantErr       bool
	}{
		{
			name: "success",
			prepareMock: func() {
				ctx := suite.ec.Request().Context()

				payrollPeriod.PayrollRunAt = &payrollPeriod.CreatedAt

				totalTakeHomePay := uint64(100000)
				expected = &response.PayslipSummary{
					TotalTakeHomePay: totalTakeHomePay,
				}

				suite.repo.EXPECT().FindPayrollPeriodByID(ctx, gomock.Eq(payrollPeriodID)).
					Return(payrollPeriod, nil).
					Times(1)
				suite.repo.EXPECT().FindPayslipSumTotalSalary(ctx, gomock.Eq(payrollPeriodID)).
					Return(totalTakeHomePay, nil).
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
			result, err := suite.payslipService.GetSummary(suite.ec, payrollPeriodID)

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
