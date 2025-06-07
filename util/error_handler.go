package util

import (
	"errors"
	"fmt"
	"github.com/cchristian77/payroll_be/response"
	sharedErrs "github.com/cchristian77/payroll_be/util/errors"
	"github.com/cchristian77/payroll_be/util/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// ErrorHandler returns JSON including status code and error message if error occurs
func ErrorHandler(err error, ec echo.Context) {
	var statusCode int
	var errorMsg string

	// Get status code and error message from if error is HTTP error type
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		statusCode = httpError.Code
		errorMsg = fmt.Sprintf("%s", httpError.Message)
	} else {
		statusCode = getStatusCode(err)
		errorMsg = err.Error()
	}

	// record error to log
	logger.Error(errorMsg)

	// Return JSON with status code and error message
	if !ec.Response().Committed {
		ec.JSON(statusCode, response.Error{Message: errorMsg, Status: statusCode})
	}
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	var businessValidationErr sharedErrs.BusinessValidationErr

	switch {
	case errors.Is(err, sharedErrs.InternalServerErr):
		return http.StatusInternalServerError
	case errors.Is(err, sharedErrs.NotFoundErr), errors.Is(err, gorm.ErrRecordNotFound):
		return http.StatusNotFound
	case errors.Is(err, sharedErrs.ConflictErr):
		return http.StatusConflict
	case errors.Is(err, sharedErrs.BadParamInputErr), errors.Is(err, sharedErrs.IncorrectCredentialErr), errors.As(err, &businessValidationErr):
		return http.StatusBadRequest
	case errors.Is(err, sharedErrs.ForbiddenErr):
		return http.StatusForbidden
	case errors.Is(err, sharedErrs.UnauthorizedErr), errors.Is(err, sharedErrs.InvalidTokenErr), errors.Is(err, sharedErrs.ExpiredTokenErr):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
