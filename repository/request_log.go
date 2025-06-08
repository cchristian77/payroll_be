package repository

import (
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util/logger"
	"github.com/labstack/echo/v4"
)

func (r *repo) CreateRequestLog(ec echo.Context, data *domain.RequestLog) (*domain.RequestLog, error) {
	ctx := ec.Request().Context()

	db, _ := database.ConnFromContext(ctx, r.DB)

	if err := db.WithContext(ctx).Create(&data).Error; err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on create request log : %v", err))

		return nil, err
	}

	return data, nil
}
