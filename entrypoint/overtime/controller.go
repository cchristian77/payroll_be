package overtime

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/service/overtime"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller manages overtime operations.
type Controller struct {
	overtime overtime.Service
}

// NewController initializes a new Controller instance.
func NewController(overtime overtime.Service) *Controller {
	return &Controller{overtime: overtime}
}

// RegisterRoutes configures the routes for the Controller.
func (c *Controller) RegisterRoutes(router *echo.Echo) {
	groupV1 := router.Group("/overtimes/v1", middleware.GetAuthorization().Authenticate())
	groupV1.POST("", c.Upsert)
}

func (c *Controller) Upsert(ec echo.Context) error {
	ctx := ec.Request().Context()

	var input request.UpsertOvertime

	if err := ec.Bind(&input); err != nil {
		return response.NewErrorResponse(ec, http.StatusUnprocessableEntity, "Invalid request body", err)
	}

	if err := ec.Validate(input); err != nil {
		return err
	}

	data, err := c.overtime.Upsert(ctx, &input)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}
