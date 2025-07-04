package attendance

import (
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/service/attendance"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller manages attendance operations, such as check-in and check-out.
type Controller struct {
	attendance attendance.Service
}

// NewController initializes a new Controller instance.
func NewController(attendance attendance.Service) *Controller {
	return &Controller{attendance: attendance}
}

// RegisterRoutes configures the routes for the Controller.
func (c *Controller) RegisterRoutes(router *echo.Echo) {
	groupV1 := router.Group("/attendances/v1", middleware.GetAuthorization().Authenticate())
	groupV1.POST("/check_in", c.CheckIn)
	groupV1.POST("/check_out", c.CheckOut)
}

func (c *Controller) CheckIn(ec echo.Context) error {
	ctx := ec.Request().Context()

	data, err := c.attendance.CheckIn(ctx)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}

func (c *Controller) CheckOut(ec echo.Context) error {
	ctx := ec.Request().Context()

	data, err := c.attendance.CheckOut(ctx)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}
