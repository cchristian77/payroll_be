package auth

import (
	"context"
	"errors"
	"github.com/cchristian77/payroll_be/domain"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/cchristian77/payroll_be/util/token"
	"gorm.io/gorm"
)

func (b base) Authenticate(ctx context.Context, accessToken string) (*domain.User, error) {
	payload, err := token.Get().Verify(accessToken)
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
