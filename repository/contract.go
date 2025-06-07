package repository

import (
	"context"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/util"
	"time"
)

//go:generate mockgen -package repository -source=contract.go -destination=mock_repository.go *

type Repository interface {

	// Session
	CreateSession(ctx context.Context, data *domain.Session) (*domain.Session, error)
	FindSessionByID(ctx context.Context, id uint64) (*domain.Session, error)
	DeleteSessionByID(ctx context.Context, id uint64) error
	RevokeSessionByID(ctx context.Context, id uint64) error

	// User
	FindUserByUsername(ctx context.Context, username string) (*domain.User, error)
	FindUserByID(ctx context.Context, id uint64) (*domain.User, error)
	FindBatchUsers(ctx context.Context, batchSize int, lastID uint64) ([]*domain.User, error)
	CreateUser(ctx context.Context, data *domain.User) (*domain.User, error)

	// Attendance
	FindAttendanceByUserIDAndDate(ctx context.Context, userID uint64, date time.Time) (*domain.Attendance, error)
	FindAttendanceByIDAndUserID(ctx context.Context, id, userID uint64) (*domain.Attendance, error)
	FindAttendancesByUserIDAndDateRange(ctx context.Context, userID uint64, startDate, endDate time.Time) ([]*domain.Attendance, error)
	CreateAttendance(ctx context.Context, data *domain.Attendance) (*domain.Attendance, error)
	UpdateAttendance(ctx context.Context, data *domain.Attendance) error

	// Overtime
	FindOvertimeByUserIDAndDate(ctx context.Context, userID uint64, date time.Time) (*domain.Overtime, error)
	FindOvertimeByIDAndUserID(ctx context.Context, id, userID uint64) (*domain.Overtime, error)
	FindOvertimesByUserIDAndDateRange(ctx context.Context, userID uint64, startDate, endDate time.Time) ([]*domain.Overtime, error)
	UpsertOvertime(ctx context.Context, data *domain.Overtime) (*domain.Overtime, error)

	// Reimbursement
	FindReimbursementByIDAndUserID(ctx context.Context, id, userID uint64) (*domain.Reimbursement, error)
	FindReimbursementsByUserIDAndStatus(ctx context.Context, userID uint64, status string) ([]*domain.Reimbursement, error)
	FindReimbursementsByPayslipID(ctx context.Context, payslipID uint64) ([]*domain.Reimbursement, error)
	UpsertReimbursement(ctx context.Context, data *domain.Reimbursement) (*domain.Reimbursement, error)

	// Payroll Period
	FindPayrollPeriodByID(ctx context.Context, id uint64) (*domain.PayrollPeriod, error)
	FindOverlappingPayrollPeriods(ctx context.Context, startDate, endDate time.Time) ([]domain.PayrollPeriod, error)
	UpsertPayrollPeriod(ctx context.Context, data *domain.PayrollPeriod) (*domain.PayrollPeriod, error)

	// Payslip
	FindPayslipByUserIDAndPayrollPeriodID(ctx context.Context, userID, payrollPeriodID uint64) (*domain.Payslip, error)
	FindPayslipPaginated(ctx context.Context, payrollPeriodID uint64, search string, p *util.Pagination) ([]*domain.Payslip, error)
	FindPayslipSumTotalSalary(ctx context.Context, payrollPeriodID uint64) (uint64, error)
	CreatePayslip(ctx context.Context, data *domain.Payslip) (*domain.Payslip, error)
}
