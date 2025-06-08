package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util/logger"
)

func (r *repo) CreateRequestLog(ctx context.Context, data *domain.RequestLog) (*domain.RequestLog, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		logger.Error(ctx, fmt.Sprintf("[REPOSITORY] Failed on create request log : %v", err))

		return nil, err
	}

	return data, nil
}
