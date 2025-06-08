package payslip

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/util"
	"github.com/labstack/echo/v4"
)

// FindPayslipList retrieves the paginated list of payslip for each employee.
func (b *base) FindPayslipList(ec echo.Context, input *request.FindPayslipList) (*response.BasePagination[[]*response.Payslip], error) {
	ctx := ec.Request().Context()

	var p util.Pagination
	p.SetPage(input.Page)
	p.SetLimit(input.PerPage)

	payrollPeriod, err := b.EnsurePayrollExecuted(ec, input.PayrollPeriodID)
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
