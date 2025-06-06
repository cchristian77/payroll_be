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

func (r *repo) FindAttendanceByUserIDAndDate(ctx context.Context, userID uint64, date time.Time) (*domain.Attendance, error) {
	var data *domain.Attendance

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("user_id = ? AND date = ?", userID, date.Format("2006-01-02")).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find attendance by user id and date : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindAttendanceByIDAndUserID(ctx context.Context, id, userID uint64) (*domain.Attendance, error) {
	var data *domain.Attendance

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find attendance by id and user id: %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) CreateAttendance(ctx context.Context, data *domain.Attendance) (*domain.Attendance, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on create attendance : %v", err))

		return data, err
	}

	return data, nil
}

func (r *repo) UpdateAttendance(ctx context.Context, data *domain.Attendance) error {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on update attendance : %v", err))

		return err
	}

	return nil
}
