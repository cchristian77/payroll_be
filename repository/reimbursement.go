package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/gorm/clause"
)

func (r *repo) FindReimbursementByIDAndUserID(ctx context.Context, id, userID uint64) (*domain.Reimbursement, error) {
	var data *domain.Reimbursement

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find reimbursement by id and user id : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindReimbursementsByUserIDAndStatus(ctx context.Context, userID uint64, status string) ([]*domain.Reimbursement, error) {
	var data []*domain.Reimbursement

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("user_id = ? AND status = ?", userID, status).
		Find(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find reimbursements by id and status : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindReimbursementsByPayslipID(ctx context.Context, payslipID uint64) ([]*domain.Reimbursement, error) {

	var data []*domain.Reimbursement

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("payslip_id = ? AND status = ?", payslipID, enums.PAIDReimbursementStatus).
		Find(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find reimbursements by payslip id : %v", err))

		return nil, err
	}

	return data, nil

}

func (r *repo) UpsertReimbursement(ctx context.Context, data *domain.Reimbursement) (*domain.Reimbursement, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Clauses(
			clause.Returning{},
			clause.OnConflict{
				Columns: []clause.Column{{Name: "id"}},
				DoUpdates: clause.Assignments(map[string]any{
					"updated_at":  data.UpdatedAt,
					"updated_by":  data.UpdatedBy,
					"description": data.Description,
					"amount":      data.Amount,
				}),
			}).
		Omit("updated_by").
		Create(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on upsert reimbursement : %v", err))

		return data, err
	}

	return data, nil
}
