package reimbursement

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/service/reimbursement"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller manages reimbursement operations.
type Controller struct {
	reimbursement reimbursement.Service
}

// NewController initializes a new Controller instance.
func NewController(reimbursement reimbursement.Service) *Controller {
	return &Controller{reimbursement: reimbursement}
}

// RegisterRoutes configures the routes for the Controller.
func (c *Controller) RegisterRoutes(router *echo.Echo) {
	groupV1 := router.Group("/reimbursements/v1", middleware.GetAuthorization().Authenticate())
	groupV1.POST("", c.Upsert)
}

func (c *Controller) Upsert(ec echo.Context) error {
	ctx := ec.Request().Context()

	var input request.UpsertReimbursement

	if err := ec.Bind(&input); err != nil {
		return response.NewErrorResponse(ec, http.StatusUnprocessableEntity, "Invalid request body", err)
	}

	if err := ec.Validate(input); err != nil {
		return err
	}

	data, err := c.reimbursement.Upsert(ctx, &input)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}
