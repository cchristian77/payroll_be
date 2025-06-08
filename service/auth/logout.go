package auth

import (
	"context"
	"errors"
	"github.com/cchristian77/payroll_be/domain/enums"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"gorm.io/gorm"
)

func (b *base) Logout(ctx context.Context) error {
	sessionID, ok := ctx.Value(enums.SessionIDCtxKey).(string)
	if !ok {
		return sharedErrs.InvalidTokenErr
	}

	session, err := b.repository.FindSessionBySessionID(ctx, sessionID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if session == nil {
		return sharedErrs.InvalidTokenErr
	}

	if err = b.repository.DeleteSessionByID(ctx, session.ID); err != nil {
		return err
	}

	return nil
}
