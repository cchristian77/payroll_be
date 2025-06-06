package response

import (
	"github.com/labstack/echo/v4"
)

type Success struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   error  `json:"error,omitempty"`
}

func NewErrorResponse(ec echo.Context, statusCode int, message string, err error) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ec.Response().WriteHeader(statusCode)

	return ec.JSON(statusCode, Error{
		Message: err.Error(),
		Status:  statusCode,
		Error:   err,
	})
}

func NewSuccessResponse(ec echo.Context, statusCode int, data any) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ec.Response().WriteHeader(statusCode)

	return ec.JSON(statusCode, Success{
		Message: "OK",
		Data:    data,
	})
}

func NewSuccessMessageResponse(ec echo.Context, statusCode int, message string) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ec.Response().WriteHeader(statusCode)

	return ec.JSON(statusCode, Success{
		Message: message,
	})
}
