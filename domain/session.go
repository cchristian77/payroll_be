package domain

import "time"

type Session struct {
	ID                   uint64
	UserID               uint64
	AccessToken          string
	AccessTokenExpiresAt time.Time
	AccessTokenCreatedAt time.Time
	UserAgent            string
	ClientIP             string
	IsRevoked            bool
}
