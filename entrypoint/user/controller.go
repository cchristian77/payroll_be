package user

import (
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/service/payslip"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Controller manage handler for user or employee operations.
type Controller struct {
	payslip payslip.Service
}

// NewController initializes a new Controller instance.
func NewController(payslip payslip.Service) *Controller {
	return &Controller{
		payslip: payslip,
	}
}

// RegisterRoutes configures the routes for the Controller.
func (c *Controller) RegisterRoutes(router *echo.Echo) {
	groupV1 := router.Group("/user/v1", middleware.GetAuthorization().Authenticate())

	payrollGroup := groupV1.Group("/payslips")
	payrollGroup.GET("", c.MyPayslip)
}

// MyPayslip retrieves the payslip on the specified payroll_period_id for the authenticated user.
func (c *Controller) MyPayslip(ec echo.Context) error {
	ctx := ec.Request().Context()

	payrollPeriodID, err := strconv.Atoi(ec.QueryParam("payroll_period_id"))
	if err != nil || payrollPeriodID <= 0 {
		return response.NewErrorResponse(ec, http.StatusBadRequest, "Please provide a valid payroll_period_id as integer", err)
	}

	result, err := c.payslip.FindUserPayslip(ctx, uint64(payrollPeriodID))
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, result)
}
