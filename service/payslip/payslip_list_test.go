package payslip

import (
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	"github.com/cchristian77/payroll_be/util/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func (suite *PayslipTestSuite) Test_PayslipList() {
	payrollPeriod := mock.InitPayrollPeriodDomain()
	payslips := make([]*domain.Payslip, 0)

	input := &request.FindPayslipList{
		Page:            1,
		PerPage:         10,
		PayrollPeriodID: 1,
	}

	var expected []*response.Payslip

	for i := 0; i < 5; i++ {
		p := mock.InitPayslipDomain()
		p.ID = uint64(i + 1)
		p.PayrollPeriod = payrollPeriod
		payslips = append(payslips, p)
	}

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
				for _, payslip := range payslips {
					expected = append(expected, response.NewPayslipFromDomain(payslip))
				}

				suite.repo.EXPECT().FindPayrollPeriodByID(ctx, gomock.Eq(input.PayrollPeriodID)).
					Return(payrollPeriod, nil).
					Times(1)
				suite.repo.EXPECT().FindPayslipPaginated(ctx, gomock.Eq(input.PayrollPeriodID), gomock.Eq(input.Search), gomock.Any()).
					Return(payslips, nil).
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
			result, err := suite.payslipService.FindPayslipList(suite.ec, input)

			// Assert
			assert.Equal(t, tc.wantErr, err != nil, "error expected %v, but actual: %v", tc.wantErr, err)
			if tc.wantErr {
				assert.Empty(t, result)
				assert.Error(t, err)
			} else {
				assert.NotEmpty(t, result)
				for i, actual := range result.Data {
					if err = util.CompareData(actual, expected[i], 1); err != nil {
						t.Fatalf("error on comparing data : %v", err)
					}
				}

			}
		})
	}
}
