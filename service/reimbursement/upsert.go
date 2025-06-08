package reimbursement

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"time"
)

func (b *base) Upsert(ctx context.Context, input *request.UpsertReimbursement) (*response.Reimbursement, error) {
	authUser := util.AuthUserFromCtx(ctx)

	now := time.Now()

	if input.ID != 0 {
		// Check whether the reimbursement exists on update
		reimbursementExists, err := b.repository.FindReimbursementByIDAndUserID(ctx, input.ID, authUser.ID)
		if err != nil {
			return nil, err
		}

		if reimbursementExists == nil {
			return nil, sharedErrs.NotFoundErr
		}

		// paid reimbursement cannot be updated.
		if reimbursementExists.Status == enums.PAIDReimbursementStatus {
			return nil, sharedErrs.NewBusinessValidationErr("Reimbursement has already been paid.")
		}
	}

	reimbursement, err := b.repository.UpsertReimbursement(ctx, &domain.Reimbursement{
		BaseModel: domain.BaseModel{
			ID:        input.ID,
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: authUser.ID,
			UpdatedBy: &authUser.ID,
		},
		UserID:      authUser.ID,
		Description: input.Description,
		Amount:      input.Amount,
		Status:      enums.PENDINGReimbursementStatus,
	})
	if err != nil {
		return nil, err
	}

	return response.NewReimbursementFromDomain(reimbursement), nil
}
