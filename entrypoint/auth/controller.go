package auth

import (
	"github.com/cchristian77/payroll_be/request"
	"github.com/cchristian77/payroll_be/response"
	"github.com/cchristian77/payroll_be/service/auth"
	"github.com/cchristian77/payroll_be/util"
	"github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller manages the authentication operations, such as login, logout, etc.
type Controller struct {
	auth auth.Service
}

// NewController initializes a new Controller instance.
func NewController(auth auth.Service) *Controller {
	return &Controller{auth: auth}
}

func (c *Controller) RegisterRoutes(router *echo.Echo) {
	groupV1 := router.Group("/auth/v1")
	groupV1.POST("/login", c.Login)
	groupV1.POST("/logout", c.Logout, middleware.GetAuthorization().Authenticate())
	groupV1.GET("/me", c.CurrentUser, middleware.GetAuthorization().Authenticate())
	groupV1.POST("/register", c.Register)
}

func (c *Controller) Login(ec echo.Context) error {
	var input request.Login

	if err := ec.Bind(&input); err != nil {
		return response.NewErrorResponse(ec, http.StatusUnprocessableEntity, "Invalid request body", err)
	}

	if err := ec.Validate(input); err != nil {
		return err
	}

	data, err := c.auth.Login(ec, &input)
	if err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, data)
}

func (c *Controller) Logout(ec echo.Context) error {
	if err := c.auth.Logout(ec); err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, "Logout success.")
}

func (c *Controller) CurrentUser(ec echo.Context) error {
	authUser := util.EchoCntextAuthUser(ec)

	return response.NewSuccessResponse(ec, http.StatusOK, response.User{
		ID:       authUser.ID,
		Username: authUser.Username,
		FullName: authUser.FullName,
		Role:     authUser.Role,
	})
}

func (c *Controller) Register(ec echo.Context) error {
	if err := c.auth.Register(ec); err != nil {
		return err
	}

	return response.NewSuccessResponse(ec, http.StatusOK, nil)
}
