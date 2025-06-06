package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/gorm/clause"
	"time"
)

func (r *repo) FindAttendanceByUserIDAndDate(ctx context.Context, userID uint64, date time.Time) (*domain.Attendance, error) {
	var data *domain.Attendance

	err := r.DB.WithContext(ctx).
		Where("user_id = ? AND date = ?", userID, date.Format("2006-01-02")).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find attendance exists by user id and date : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindAttendanceByID(ctx context.Context, id uint64) (*domain.Attendance, error) {
	var data *domain.Attendance

	err := r.DB.WithContext(ctx).
		First(&data, id).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find attendance by id : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) CreateAttendance(ctx context.Context, data *domain.Attendance) (*domain.Attendance, error) {
	err := r.DB.WithContext(ctx).
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
	err := r.DB.WithContext(ctx).
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
