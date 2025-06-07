package auth

import (
	"errors"
	"github.com/cchristian77/payroll_be/domain/enums"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (b *base) Logout(ec echo.Context) error {
	ctx := ec.Request().Context()

	sessionID, ok := ec.Get(enums.SessionIDCtxKey).(string)
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
