package payslip

import (
	"context"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
)

// FindPayslipList retrieves the paginated list of payslip for each employee.
func (b *base) FindPayslipList(ctx context.Context, input *request.FindPayslipList) (*response.BasePagination[[]*response.Payslip], error) {

	var p util.Pagination
	p.SetPage(input.Page)
	p.SetLimit(input.PerPage)

	payrollPeriod, err := b.EnsurePayrollExecuted(ctx, input.PayrollPeriodID)
	if err != nil {
		return nil, err
	}

	payslips, err := b.repository.FindPayslipPaginated(ctx, input.PayrollPeriodID, input.Search, &p)
	if err != nil {
		return nil, err
	}

	result := make([]*response.Payslip, len(payslips))
	for i, payslip := range payslips {
		payslip.PayrollPeriod = payrollPeriod

		result[i] = response.NewPayslipFromDomain(payslip)
	}

	return &response.BasePagination[[]*response.Payslip]{
		Data: result,
		Meta: &response.Meta{
			Page:      p.Page(),
			PerPage:   len(result),
			PageCount: p.PageCount(),
			Total:     p.Total(),
		},
	}, nil
}
