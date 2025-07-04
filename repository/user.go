package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util/logger"
)

func (r *repo) FindUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var result *domain.User

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("username = ?", username).
		First(&result).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on find user by username : %v", err))

		return nil, err
	}

	return result, nil
}

func (r *repo) FindUserByID(ctx context.Context, id uint64) (*domain.User, error) {
	var result *domain.User

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		First(&result, id).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on find user by id : %v", err))

		return nil, err
	}

	return result, nil
}

func (r *repo) FindBatchUsers(ctx context.Context, batchSize int, lastID uint64) ([]*domain.User, error) {
	var result []*domain.User

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("id > ?", lastID).
		Limit(batchSize).
		Find(&result).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on find batch users : %v", err))

		return nil, err
	}

	return result, nil
}

func (r *repo) CreateUser(ctx context.Context, data *domain.User) (*domain.User, error) {
	var result *domain.User

	err := r.DB.WithContext(ctx).
		Create(&data).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on create user : %v", err))

		return nil, err
	}

	return result, nil

}
