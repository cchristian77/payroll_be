package logger

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
)

func authUserFromContext(ctx context.Context) *domain.User {
	authUser, ok := ctx.Value(enums.AuthUserCtxKey).(*domain.User)
	if !ok {
		return nil
	}

	return authUser
}

func requestIDFromContext(ctx context.Context) string {
	requestID, ok := ctx.Value(enums.RequestIDCtxKey).(string)
	if !ok {
		return ""
	}

	return requestID
}
