package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/gorm/clause"
)

func (r *repo) CreateSession(ctx context.Context, data *domain.Session) (*domain.Session, error) {
	err := r.DB.WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on create session : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindSessionByID(ctx context.Context, id uint64) (*domain.Session, error) {
	var result *domain.Session

	err := r.DB.WithContext(ctx).
		First(&result, id).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find session by id : %v", err))

		return nil, err
	}

	return result, nil
}

func (r *repo) DeleteSessionByID(ctx context.Context, id uint64) error {
	err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Session{}).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on delete session by id : %v", err))

		return err
	}

	return nil
}

func (r *repo) RevokeSessionByID(ctx context.Context, id uint64) error {
	err := r.DB.
		WithContext(ctx).
		Model(&domain.Session{}).
		Where("id = ?", id).
		Update("is_revoked", true).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on revoke session by id : %v", err))

		return err
	}

	return nil
}
