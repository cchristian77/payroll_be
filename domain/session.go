package domain

import "time"

type Session struct {
	ID                    uint64
	UserID                uint64
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	AccessTokenCreatedAt  time.Time
	RefreshTokenExpiresAt time.Time
	RefreshTokenCreatedAt time.Time
	UserAgent             string
	ClientIP              string
	IsRevoked             bool
}
