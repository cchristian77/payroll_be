package auth

import (
	"errors"
	"github.com/cchristian77/payroll_be/domain"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	tokenMaker "github.com/cchristian77/payroll_be/util/token"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (b *base) Authenticate(ec echo.Context, accessToken string) (*domain.User, error) {
	ctx := ec.Request().Context()

	payload, err := tokenMaker.Get().Verify(accessToken)
	if err != nil {
		return nil, err
	}

	authUser, err := b.repository.FindUserByID(ctx, payload.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sharedErrs.IncorrectCredentialErr
		}

		return nil, err
	}

	return authUser, nil
}
