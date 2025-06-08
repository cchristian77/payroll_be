package reimbursement

import (
	"github.com/cchristian77/payroll_be/domain/enums"
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

func (suite *ReimbursementServiceTestSuite) Test_Upsert() {
	reimbursement := mock.InitReimbursementDomain()
	input := &request.UpsertReimbursement{
		Description: "test_description",
		Amount:      5000,
	}

	var expected *response.Reimbursement

	testCases := []struct {
		name          string
		prepareMock   func()
		expectedError error
		wantErr       bool
	}{
		{
			name: "reimbursement not found on update",
			prepareMock: func() {
				authUser := util.AuthUserFromCtx(suite.ctx)

				input.ID = 1

				suite.repo.EXPECT().FindReimbursementByIDAndUserID(suite.ctx, gomock.Eq(input.ID), gomock.Eq(authUser.ID)).
					Return(nil, gorm.ErrRecordNotFound).
					Times(1)
				suite.repo.EXPECT().UpsertReimbursement(suite.ctx, gomock.Any()).Times(0)
			},
			wantErr:       true,
			expectedError: gorm.ErrRecordNotFound,
		},
		{
			name: "reimbursement is already paid",
			prepareMock: func() {
				authUser := util.AuthUserFromCtx(suite.ctx)

				input.ID = 1
				reimbursement.Status = enums.PAIDReimbursementStatus

				suite.repo.EXPECT().FindReimbursementByIDAndUserID(suite.ctx, gomock.Eq(input.ID), gomock.Eq(authUser.ID)).
					Return(reimbursement, nil).
					Times(1)
				suite.repo.EXPECT().UpsertReimbursement(suite.ctx, gomock.Any()).Times(0)
			},
			wantErr:       true,
			expectedError: sharedErrs.NewBusinessValidationErr("Reimbursement has already been paid."),
		},
		{
			name: "success",
			prepareMock: func() {
				authUser := util.AuthUserFromCtx(suite.ctx)

				reimbursement.Status = enums.PENDINGReimbursementStatus
				expected = response.NewReimbursementFromDomain(reimbursement)

				suite.repo.EXPECT().FindReimbursementByIDAndUserID(suite.ctx, gomock.Eq(authUser.ID), gomock.Any()).
					Return(reimbursement, nil).
					Times(1)
				suite.repo.EXPECT().UpsertReimbursement(suite.ctx, gomock.Any()).
					Return(reimbursement, nil).
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
			result, err := suite.reimbursementService.Upsert(suite.ctx, input)

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
