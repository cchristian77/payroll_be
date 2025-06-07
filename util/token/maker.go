package token

import (
	"github.com/google/uuid"
	"time"
)

// Maker is an interface for managing tokens
type Maker interface {
	// GenerateToken creates a new token for a specific username and duration
	Generate(sessionID uuid.UUID, userID uint64, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	Verify(token string) (*Payload, error)
}
