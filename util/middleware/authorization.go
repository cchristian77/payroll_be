package middleware

import (
	"fmt"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/service/auth"
	"github.com/cchristian77/payroll_be/util"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// authHeaderKey represents the key for the Authorization header in HTTP requests.
// authTypeBearer defines the bearer authentication type used in the Authorization header.
const (
	authHeaderKey  = "authorization"
	authTypeBearer = "bearer"
)

var authMiddleware *Authorization

// Authorization is a middleware model that provides role-based authorization for HTTP requests.
type Authorization struct {
	authService auth.Service
}

// InitAuthorization initializes the authorization middleware with the provided authentication service.
func InitAuthorization(authService auth.Service) {
	if authMiddleware == nil {
		authMiddleware = &Authorization{authService: authService}
	}

	return
}

// GetAuthorization returns a pointer to the Authorization instance initialized as middleware.
func GetAuthorization() *Authorization {
	return authMiddleware
}

// AdminOnly is an HTTP middleware that restricts access only to users with the ADMIN role.
func (a *Authorization) AdminOnly() echo.MiddlewareFunc {
	return a.authenticationWithRoles(enums.ADMINRole)
}

// UserOnly is a middleware function that restricts access only to users with the USER role.
func (a *Authorization) UserOnly() echo.MiddlewareFunc {
	return a.authenticationWithRoles(enums.USERRole)
}

// Authenticate provides middleware for authorizing requests by validating the user's role as either ADMIN or USER.
func (a *Authorization) Authenticate() echo.MiddlewareFunc {
	return a.authenticationWithRoles(enums.ADMINRole, enums.USERRole)
}

// authenticationWithRoles creates middleware to authenticate users and check if they have one of the specified roles.
func (a *Authorization) authenticationWithRoles(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			/// gather token from the header
			authHeader := ec.Request().Header.Get(authHeaderKey)
			if authHeader == "" {
				return sharedErrs.UnauthorizedErr
			}

			authFields := strings.Fields(authHeader)
			if len(authFields) < 2 {
				return sharedErrs.UnauthorizedErr
			}

			authorizationType := strings.ToLower(authFields[0])
			if authorizationType != authTypeBearer {
				return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("unsupported authorization type %s", authorizationType))
			}

			bearerToken := authFields[1]

			// authenticate the bearer token
			authUser, payload, err := a.authService.Authenticate(ec, bearerToken)
			if err != nil {
				return err
			}

			// authorization based on the allowed roles from the authenticated user's role
			// find if user's role contains in the allowed roles
			if util.Contains(allowedRoles, authUser.Role) {
				ec.Set(enums.AuthUserCtxKey, authUser)
				ec.Set(enums.SessionIDCtxKey, payload.ID.String())

				return next(ec)
			}

			// if not found, then return forbideen access error
			return sharedErrs.ForbiddenErr
		}
	}
}
