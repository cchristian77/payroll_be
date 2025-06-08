package request_log

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/util"
	"time"
)

func (b *base) Log(ctx context.Context, activity string, referenceID uint64, entity string) (*domain.RequestLog, error) {
	authUser := util.AuthUserFromCtx(ctx)

	now := time.Now()
	requestLog, err := b.repository.CreateRequestLog(ctx, &domain.RequestLog{
		CreatedAt:   now,
		UpdatedAt:   now,
		RequestID:   ctx.Value(enums.RequestIDCtxKey).(string),
		UserID:      authUser.ID,
		Activity:    activity,
		Entity:      entity,
		ReferenceID: referenceID,
		ClientIP:    ctx.Value(enums.IPAddressCtxKey).(string),
	})
	if err != nil {
		return nil, err
	}

	return requestLog, nil
}
