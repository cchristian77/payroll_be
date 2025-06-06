package controller

import (
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/service/auth"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	auth auth.Service
}

func NewController(auth auth.Service) *Controller {
	return &Controller{auth: auth}
}

func (c *Controller) RegisterRoutes(router *echo.Echo) {
	group := router.Group("/auth")
	group.POST("/login", c.Login)
	group.GET("/current_user", c.CurrentUser, middleware.GetAuthorization().Authenticate())
}

func (c *Controller) Login(ec echo.Context) error {
	ctx := ec.Request().Context()

	var input request.Login

	if err := ec.Bind(&input); err != nil {
		return response.NewErrorResponse(ec, http.StatusBadRequest, "Invalid request body", err)
	}

	if err := ec.Validate(input); err != nil {
		return err
	}

	input.UserAgent = ec.Request().UserAgent()
	input.IPAddress = ec.RealIP()

	data, err := c.auth.Login(ctx, &input)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}

func (c *Controller) CurrentUser(ec echo.Context) error {
	authUser := ec.Get(middleware.AuthUserKey).(*domain.User)
	return response.NewSuccessResponse(ec, http.StatusOK, response.User{
		ID:       authUser.ID,
		Username: authUser.Username,
		FullName: authUser.FullName,
		Role:     authUser.Role,
	})
}
