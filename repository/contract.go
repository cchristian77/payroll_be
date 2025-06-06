package repository

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"time"
)

type Repository interface {

	// Session
	CreateSession(ctx context.Context, data *domain.Session) (*domain.Session, error)
	FindSessionByID(ctx context.Context, id uint64) (*domain.Session, error)
	DeleteSessionByID(ctx context.Context, id uint64) error
	RevokeSessionByID(ctx context.Context, id uint64) error

	// User
	FindUserByUsername(ctx context.Context, username string) (*domain.User, error)
	FindUserByID(ctx context.Context, id uint64) (*domain.User, error)

	// Attendance
	FindAttendanceByUserIDAndDate(ctx context.Context, userID uint64, date time.Time) (*domain.Attendance, error)
	FindAttendanceByID(ctx context.Context, id uint64) (*domain.Attendance, error)
	CreateAttendance(ctx context.Context, data *domain.Attendance) (*domain.Attendance, error)
	UpdateAttendance(ctx context.Context, data *domain.Attendance) error
}
