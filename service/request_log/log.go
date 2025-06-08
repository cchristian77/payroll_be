package request_log

import (
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/util"
	"github.com/labstack/echo/v4"
	"time"
)

func (b *base) Log(ec echo.Context, activity string, referenceID uint64, entity string) (*domain.RequestLog, error) {
	authUser := util.EchoCntextAuthUser(ec)

	now := time.Now()
	requestLog, err := b.repository.CreateRequestLog(ec, &domain.RequestLog{
		CreatedAt:   now,
		UpdatedAt:   now,
		RequestID:   ec.Get(enums.RequestIDCtxKey).(string),
		UserID:      authUser.ID,
		Activity:    activity,
		Entity:      entity,
		ReferenceID: referenceID,
		ClientIP:    ec.RealIP(),
	})
	if err != nil {
		return nil, err
	}

	return requestLog, nil
}
