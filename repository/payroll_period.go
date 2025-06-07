package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/gorm/clause"
	"time"
)

func (r *repo) FindOverlappingPayrollPeriods(ctx context.Context, startDate, endDate time.Time) ([]domain.PayrollPeriod, error) {
	var data []domain.PayrollPeriod

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("start_date <= ? AND end_date >= ?", startDate.Format(time.DateOnly), endDate.Format(time.DateOnly)).
		Find(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find overlapping payroll periods : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindPayrollPeriodByID(ctx context.Context, id uint64) (*domain.PayrollPeriod, error) {
	var data *domain.PayrollPeriod

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		First(&data, id).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find payroll period by id : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) UpsertPayrollPeriod(ctx context.Context, data *domain.PayrollPeriod) (*domain.PayrollPeriod, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Clauses(
			clause.Returning{},
			clause.OnConflict{
				Columns: []clause.Column{{Name: "id"}},
				DoUpdates: clause.Assignments(map[string]any{
					"updated_at":     data.UpdatedAt,
					"updated_by":     data.UpdatedBy,
					"start_date":     data.StartDate,
					"end_date":       data.EndDate,
					"payroll_run_at": data.PayrollRunAt,
				}),
			}).
		Omit("updated_by").
		Create(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on upsert payroll period : %v", err))

		return data, err
	}

	return data, nil
}
