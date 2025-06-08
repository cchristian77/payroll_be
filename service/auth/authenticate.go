package auth

import (
	"errors"
	"github.com/cchristian77/payroll_be/domain"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	tokenMaker "github.com/cchristian77/payroll_be/util/token"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

// Authenticate functions to decrypt and verify the provided access token.
func (b *base) Authenticate(ec echo.Context, accessToken string) (*domain.User, *tokenMaker.Payload, error) {
	ctx := ec.Request().Context()

	payload, err := tokenMaker.Get().Verify(accessToken)
	if err != nil {
		return nil, nil, err
	}

	// find session by verified payload ID
	session, err := b.repository.FindSessionBySessionID(ctx, payload.ID.String())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}

	if session == nil {
		return nil, nil, sharedErrs.InvalidTokenErr
	}

	if session.AccessToken != accessToken {
		return nil, nil, sharedErrs.InvalidTokenErr
	}

	// check whether session is expired
	if time.Now().After(session.AccessTokenExpiresAt) {
		return nil, nil, sharedErrs.ExpiredTokenErr
	}

	// find the user data from payload
	authUser, err := b.repository.FindUserByID(ctx, payload.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, sharedErrs.IncorrectCredentialErr
		}

		return nil, nil, err
	}

	return authUser, payload, nil
}
