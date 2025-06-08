package util

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
)

// AuthUserFromCtx retrieves the authenticated user from the context.
func AuthUserFromCtx(ctx context.Context) *domain.User {
	authUser, ok := ctx.Value(enums.AuthUserCtxKey).(*domain.User)
	if !ok {
		return nil
	}

	return authUser
}
