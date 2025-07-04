package admin

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	payrollPeriod "github.com/cchristian77/payroll_be/service/payroll_period"
	"github.com/cchristian77/payroll_be/service/payslip"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Controller manages payroll periods and payslip-related operations.
type Controller struct {
	payrollPeriod payrollPeriod.Service
	payslip       payslip.Service
}

// NewController initializes a new Controller instance.
func NewController(payrollPeriod payrollPeriod.Service, payslip payslip.Service) *Controller {
	return &Controller{
		payrollPeriod: payrollPeriod,
		payslip:       payslip,
	}
}

// RegisterRoutes configures the routes for the Controller.
func (c *Controller) RegisterRoutes(router *echo.Echo) {
	groupV1 := router.Group("/admin/v1", middleware.GetAuthorization().AdminOnly())

	payrollGroup := groupV1.Group("/payrolls")
	payrollGroup.POST("/periods", c.UpsertPayrollPeriod)
	payrollGroup.POST("/execute", c.RunPayroll)

	payslipGroup := groupV1.Group("/payslips")
	payslipGroup.GET("", c.FindPayslipList)
	payslipGroup.GET("/summary", c.PayslipSummary)
}

// UpsertPayrollPeriod creates or updates a payroll period.
func (c *Controller) UpsertPayrollPeriod(ec echo.Context) error {
	ctx := ec.Request().Context()

	var input request.UpsertPayrollPeriod

	if err := ec.Bind(&input); err != nil {
		return response.NewErrorResponse(ec, http.StatusUnprocessableEntity, "Invalid request body", err)
	}

	if err := ec.Validate(input); err != nil {
		return err
	}

	data, err := c.payrollPeriod.Upsert(ctx, &input)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}

// RunPayroll handles the execution of the payroll process for a specific payroll period.
func (c *Controller) RunPayroll(ec echo.Context) error {
	ctx := ec.Request().Context()

	var input request.RunPayroll

	if err := ec.Bind(&input); err != nil {
		return response.NewErrorResponse(ec, http.StatusUnprocessableEntity, "Invalid request body", err)
	}

	if err := ec.Validate(input); err != nil {
		return err
	}

	if err := c.payslip.RunPayroll(ctx, &input); err != nil {
		return err
	}

	return response.NewSuccessMessageResponse(ec, http.StatusOK, "payroll run successfully")
}

// FindPayslipList retrieves a paginated list of payslips.
func (c *Controller) FindPayslipList(ec echo.Context) error {
	ctx := ec.Request().Context()

	var input request.FindPayslipList
	var err error

	input.Page, err = strconv.Atoi(ec.QueryParam("page"))
	if err != nil || input.Page <= 0 {
		return response.NewErrorResponse(ec, http.StatusBadRequest, "Please provide a valid page as integer", err)
	}

	input.PerPage, err = strconv.Atoi(ec.QueryParam("per_page"))
	if err != nil || input.PerPage <= 0 {
		return response.NewErrorResponse(ec, http.StatusBadRequest, "Please provide a valid per_page as integer", err)
	}

	payrollPeriodID, err := strconv.Atoi(ec.QueryParam("payroll_period_id"))
	if err != nil || payrollPeriodID <= 0 {
		return response.NewErrorResponse(ec, http.StatusBadRequest, "Please provide a valid payroll_period_id as integer", err)
	}
	input.PayrollPeriodID = uint64(payrollPeriodID)
	input.Search = ec.QueryParam("search")

	result, err := c.payslip.FindPayslipList(ctx, &input)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, result)
}

// PayslipSummary fetches a summary of payslips for a specified payroll period ID.
func (c *Controller) PayslipSummary(ec echo.Context) error {
	ctx := ec.Request().Context()

	payrollPeriodID, err := strconv.Atoi(ec.QueryParam("payroll_period_id"))
	if err != nil || payrollPeriodID <= 0 {
		return response.NewErrorResponse(ec, http.StatusBadRequest, "Please provide a valid payroll_period_id as integer", err)
	}

	result, err := c.payslip.GetSummary(ctx, uint64(payrollPeriodID))
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, result)
}
