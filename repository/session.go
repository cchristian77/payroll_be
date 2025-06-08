package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/gorm/clause"
)

func (r *repo) CreateSession(ctx context.Context, data *domain.Session) (*domain.Session, error) {
	err := r.DB.WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(data).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on create session : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) FindSessionBySessionID(ctx context.Context, sessionID string) (*domain.Session, error) {
	var result *domain.Session

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("session_id = ?", sessionID).
		First(&result).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on find session by id : %v", err))

		return nil, err
	}

	return result, nil
}

func (r *repo) DeleteSessionByID(ctx context.Context, id uint64) error {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Session{}).
		Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on delete session by id : %v", err))

		return err
	}

	return nil
}
