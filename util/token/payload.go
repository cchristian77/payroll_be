package token

import (
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

// Payload contains the payload data of the token
type Payload struct {
	ID     uuid.UUID `json:"id"`
	UserID uint64    `json:"user_id"`
	jwt.StandardClaims
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(sessionID uuid.UUID, userID uint64, duration time.Duration) (*Payload, error) {
	return &Payload{
		ID:     sessionID,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "payroll.api.auth",
		},
	}, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	expiredAt := time.Unix(payload.StandardClaims.ExpiresAt, 0)

	if time.Now().After(expiredAt) {
		return sharedErrs.InvalidTokenErr
	}

	return nil
}
