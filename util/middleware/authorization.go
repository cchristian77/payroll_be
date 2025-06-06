package middleware

import (
	"fmt"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/service/auth"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authHeaderKey  = "authorization"
	authTypeBearer = "bearer"
	AuthPayloadKey = "auth_payload"
	AuthUserKey    = "auth_user"
)

var authMiddleware *Authorization

type Authorization struct {
	authService auth.Service
}

func InitAuthorization(authService auth.Service) {
	if authMiddleware == nil {
		authMiddleware = &Authorization{authService: authService}
	}

	return
}

func GetAuthorization() *Authorization {
	return authMiddleware
}

func (a *Authorization) AuthAdminOnly() echo.MiddlewareFunc {
	return a.authenticationWithRoles(enums.AdminRole)
}

func (a *Authorization) AuthUserOnly() echo.MiddlewareFunc {
	return a.authenticationWithRoles(enums.UserRole)
}

func (a *Authorization) Authenticate() echo.MiddlewareFunc {
	return a.authenticationWithRoles(enums.AdminRole, enums.UserRole)
}

func (a *Authorization) authenticationWithRoles(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			ctx := ec.Request().Context()

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

			authUser, err := a.authService.Authenticate(ctx, bearerToken)
			if err != nil {
				return err
			}

			for _, allowedRole := range allowedRoles {
				if authUser.Role == allowedRole {
					ec.Set(AuthUserKey, authUser)
					return next(ec)
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Not allowed")
		}
	}
}
