package payslip

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/cchristian77/payroll_be/util/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

func (b *base) RunPayroll(ec echo.Context, input *request.RunPayroll) error {
	ctx := ec.Request().Context()

	payrollPeriod, err := b.repository.FindPayrollPeriodByID(ctx, input.PayrollPeriodID)
	if err != nil {
		return err
	}

	if payrollPeriod.PayrollRunAt != nil {
		return sharedErrs.NewBusinessValidationErr(
			fmt.Sprintf("Payroll period %s - %s is already run at %s",
				payrollPeriod.StartDate.Format(time.DateOnly),
				payrollPeriod.EndDate.Format(time.DateOnly),
				payrollPeriod.PayrollRunAt.Format(time.DateTime)))
	}

	var (
		lastID     uint64 = 0
		batchCount        = 1
		batchSize         = 20
	)

	for {
		// Gather all users with their attendances, overtimes, reimbursements
		users, err := b.repository.FindBatchUsers(ctx, batchSize, lastID)
		if err != nil {
			return err
		}

		// break loop if no users found
		if len(users) == 0 {
			break
		}

		lastID = users[len(users)-1].ID

		for _, user := range users {
			if err = b.ProcessPayroll(ec, user, payrollPeriod); err != nil {
				return err
			}
		}

		batchCount += 1
	}

	now := time.Now()
	payrollPeriod.PayrollRunAt = &now
	_, err = b.repository.UpsertPayrollPeriod(ctx, payrollPeriod)
	if err != nil {
		return err
	}

	return nil
}

func (b *base) ProcessPayroll(ec echo.Context, user *domain.User, payrollPeriod *domain.PayrollPeriod) error {
	ctx := ec.Request().Context()

	authUser := util.EchoCntextAuthUser(ec)

	payslipExists, err := b.repository.FindPayslipByUserIDAndPayrollPeriodID(ctx, user.ID, payrollPeriod.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if payslipExists != nil {
		logger.Info(fmt.Sprintf("payslip.go %d already exists for user %d and payroll period %d",
			payslipExists.ID, user.ID, payrollPeriod.ID))
		return nil
	}

	attendancePay, totalAttendanceDays, err := b.calculateAttendancePay(ec, user, payrollPeriod)
	if err != nil {
		return err
	}

	overtimePay, totalOvertimeHours, totalOvertimeDays, err := b.calculateOvertimePay(ec, user, payrollPeriod)
	if err != nil {
		return err
	}

	reimbursementPay, err := b.calculateReimbursementPay(ec, user)
	if err != nil {
		return err
	}

	totalSalary := attendancePay + overtimePay + reimbursementPay

	tCtx, tx := database.InitTx(ctx, b.writeDB)
	defer func() {
		if err := tx.Rollback().Error; err != nil && !errors.Is(err, sql.ErrTxDone) {
			logger.Error(fmt.Sprintf("[REPOSITORY] Error on ProcessPayroll func: ROLLBACK TXN: %v", err))
		}
	}()

	// Create payslip.go for this payroll period
	now := time.Now()
	payslip, err := b.repository.CreatePayslip(tCtx, &domain.Payslip{
		BaseModel: domain.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: authUser.ID,
		},
		PayrollPeriodID:     payrollPeriod.ID,
		UserID:              user.ID,
		TotalAttendanceDays: totalAttendanceDays,
		TotalOvertimeDays:   totalOvertimeDays,
		TotalOvertimeHours:  totalOvertimeHours,
		TotalReimbursements: reimbursementPay,
		BaseSalary:          user.BaseSalary,
		AttendancePay:       attendancePay,
		OvertimePay:         overtimePay,
		ReimbursementPay:    reimbursementPay,
		TotalSalary:         totalSalary,
		PayrollPeriod:       nil,
	})
	if err != nil {
		return err
	}

	if err = b.payReimbursements(ec, user.ID, payslip.ID); err != nil {
		return err
	}

	if err = tx.Commit().Error; err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Error on RunPayroll func: COMMIT TXN: %v", err))
		return err
	}

	return nil
}

func (b *base) calculateAttendancePay(ec echo.Context, user *domain.User, payrollPeriod *domain.PayrollPeriod) (uint64, uint, error) {
	var attendancePay uint64

	attendances, err := b.repository.FindAttendancesByUserIDAndDateRange(ec.Request().Context(),
		user.ID,
		payrollPeriod.StartDate,
		payrollPeriod.EndDate)
	if err != nil {
		return 0, 0, err
	}

	totalAttendanceDays := uint64(len(attendances))

	attendancePay = totalAttendanceDays * enums.UserWorkHours * user.GetHourlyRate()

	return attendancePay, uint(totalAttendanceDays), nil
}

func (b *base) calculateOvertimePay(ec echo.Context, user *domain.User, payrollPeriod *domain.PayrollPeriod) (uint64, uint, uint, error) {
	var overtimePay uint64

	overtimes, err := b.repository.FindOvertimesByUserIDAndDateRange(ec.Request().Context(),
		user.ID,
		payrollPeriod.StartDate,
		payrollPeriod.EndDate)
	if err != nil {
		return 0, 0, 0, err
	}

	var totalOvertimeHours uint
	for _, overtime := range overtimes {
		totalOvertimeHours += overtime.Duration
	}

	overtimePay = uint64(totalOvertimeHours) * user.GetHourlyRate() * enums.UserOvertimeMultiplier

	totalOvertimeDays := uint(len(overtimes))

	return overtimePay, totalOvertimeHours, totalOvertimeDays, nil
}

func (b *base) calculateReimbursementPay(ec echo.Context, user *domain.User) (uint64, error) {
	var reimbursementPay uint64

	reimbursements, err := b.repository.FindReimbursementsByUserIDAndStatus(ec.Request().Context(),
		user.ID, enums.PENDINGReimbursementStatus)
	if err != nil {
		return 0, err
	}

	for _, reimbursement := range reimbursements {
		reimbursementPay += reimbursement.Amount
	}

	return reimbursementPay, nil
}

func (b *base) payReimbursements(ec echo.Context, userID, payslipID uint64) error {
	ctx := ec.Request().Context()

	reimbursements, err := b.repository.FindReimbursementsByUserIDAndStatus(ctx, userID, enums.PENDINGReimbursementStatus)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, reimbursement := range reimbursements {

		reimbursement.Status = enums.PAIDReimbursementStatus
		reimbursement.ReimbursedAt = &now
		reimbursement.PayslipID = &payslipID
		reimbursement.UpdatedAt = now
		if _, err := b.repository.UpsertReimbursement(ctx, reimbursement); err != nil {
			return err
		}
	}

	return nil
}
