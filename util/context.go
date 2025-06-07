package util

import (
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/labstack/echo/v4"
)

func EchoCntextAuthUser(ec echo.Context) *domain.User {
	authUser, ok := ec.Get(enums.AuthUserCtxKey).(*domain.User)
	if !ok {
		return nil
	}

	return authUser
}
