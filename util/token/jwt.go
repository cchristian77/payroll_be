package token

import (
	"github.com/cchristian77/payroll_be/util/config"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var tokenMaker Maker

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// Get creates a new JWTMaker
func Get() Maker {
	if tokenMaker == nil {
		tokenMaker = &JWTMaker{secretKey: config.Env.JWTKey}
	}

	return tokenMaker
}

// Generate creates a new token for a specific username and duration
func (maker *JWTMaker) Generate(sessionID uuid.UUID, userID uint64, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(sessionID, userID, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))

	return token, payload, err
}

// Verify checks if the token is valid or not
func (maker *JWTMaker) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, sharedErrs.InvalidTokenErr
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, sharedErrs.ExpiredTokenErr) {
			return nil, sharedErrs.ExpiredTokenErr
		}

		return nil, sharedErrs.InvalidTokenErr
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, sharedErrs.InvalidTokenErr
	}

	return payload, nil
}
