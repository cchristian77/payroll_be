package auth

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util/config"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	tokenMaker "github.com/cchristian77/payroll_be/util/token"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (b *base) Login(ctx context.Context, input *request.Login) (*response.Auth, error) {
	authUser, err := b.repository.FindUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(input.Password))
	if err != nil {
		return nil, sharedErrs.IncorrectCredentialErr
	}

	sessionID := uuid.New()
	accessTokenDuration, _ := time.ParseDuration(config.Env.Auth.AccessTokenExpiration)

	accessToken, payload, err := tokenMaker.Get().Generate(sessionID, authUser.ID, accessTokenDuration)
	if err != nil {
		return nil, err
	}

	_, err = b.repository.CreateSession(ctx, &domain.Session{
		UserID:               authUser.ID,
		SessionID:            payload.ID,
		AccessToken:          accessToken,
		AccessTokenExpiresAt: time.Unix(payload.StandardClaims.ExpiresAt, 0),
		AccessTokenCreatedAt: time.Unix(payload.StandardClaims.IssuedAt, 0),
		UserAgent:            input.UserAgent,
		ClientIP:             input.ClientIP,
	})
	if err != nil {
		return nil, err
	}

	return &response.Auth{
		User: response.User{
			ID:       authUser.ID,
			Username: authUser.Username,
			FullName: authUser.FullName,
			Role:     authUser.Role,
		},
		SessionID:            payload.ID,
		AccessToken:          accessToken,
		AccessTokenExpiresAt: payload.ExpiresAt,
	}, nil
}
