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

func (r *repo) FindOvertimeByUserIDAndDate(ctx context.Context, userID uint64, date time.Time) (*domain.Overtime, error) {
	var data *domain.Overtime

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("user_id = ? AND date = ?", userID, date.Format("2006-01-02")).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find overtime by user id and date : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindOvertimeByIDAndUserID(ctx context.Context, id, userID uint64) (*domain.Overtime, error) {
	var data *domain.Overtime

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find overtime by id and user id : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) UpsertOvertime(ctx context.Context, data *domain.Overtime) (*domain.Overtime, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Clauses(
			clause.Returning{},
			clause.OnConflict{
				Columns: []clause.Column{{Name: "attendance_id"}},
				DoUpdates: clause.Assignments(map[string]any{
					"updated_at": data.UpdatedAt,
					"updated_by": data.UpdatedBy,
					"date":       data.Date,
					"duration":   data.Duration,
				}),
			}).
		Omit("updated_by").
		Create(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on upsert overtime : %v", err))

		return data, err
	}

	return data, nil
}
