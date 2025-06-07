package repository

import (
	"context"
	"fmt"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/gorm/clause"
)

func (r *repo) FindPayslipByUserIDAndPayrollPeriodID(ctx context.Context, userID, payrollPeriodID uint64) (*domain.Payslip, error) {
	var data *domain.Payslip

	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Preload("User").
		Where("user_id = ? AND payroll_period_id = ?", userID, payrollPeriodID).
		First(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find payslip by user id and payslip.go period id : %v", err))

		return nil, err
	}

	return data, nil
}

func (r *repo) CreatePayslip(ctx context.Context, data *domain.Payslip) (*domain.Payslip, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	err := db.WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on create payslip: %v", err))

		return data, err
	}

	return data, nil
}

func (r *repo) FindPayslipSumTotalSalary(ctx context.Context, payrollPeriodID uint64) (uint64, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	var sum uint64

	err := db.WithContext(ctx).
		Model(&domain.Payslip{}).
		Select("sum(total_salary)").
		Where("payroll_period_id = ?", payrollPeriodID).
		Row().
		Scan(&sum)
	if err != nil {
		return 0, err
	}

	return sum, nil
}

func (r *repo) FindPayslipPaginated(ctx context.Context, payrollPeriodID uint64, search string, p *util.Pagination) ([]*domain.Payslip, error) {
	db, _ := database.ConnFromContext(ctx, r.DB)

	var (
		data  []*domain.Payslip
		count int64
	)

	query := db.WithContext(ctx).Model(&data).Where("payroll_period_id = ?", payrollPeriodID)

	if search != "" {
		query.Joins("JOIN users ON users.id = payslips.user_id AND users.full_name ILIKE ?", "%"+search+"%")
	}

	if err := query.Preload("User").Count(&count).Error; err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find payslip paginated count on period: %v", err))
		return nil, err
	}

	p.SetTotal(count)

	err := query.Offset(p.Offset()).
		Limit(p.Limit()).
		Find(&data).
		Error
	if err != nil {
		logger.Error(fmt.Sprintf("[REPOSITORY] Failed on find payslip.go pagionated on period : %v", err))

		return data, err
	}

	return data, nil
}
