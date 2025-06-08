package response

import (
	"github.com/labstack/echo/v4"
)

// Success represents a generic structure for successful response, containing a message and optional data.
type Success struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

// Error represents a generic structure for error response, containing a message, status code, and optional error details.
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   error  `json:"error,omitempty"`
}

// Meta represents metadata for paginated responses.
type Meta struct {
	Page      int   `json:"page,omitempty"`
	PerPage   int   `json:"per_page,omitempty"`
	PageCount int   `json:"page_count"`
	Total     int64 `json:"total"`
}

// BasePagination represents a generic paginated response structure containing data and pagination metadata.
type BasePagination[T any] struct {
	Data T     `json:"data"`
	Meta *Meta `json:"meta"`
}

// NewErrorResponse sends a JSON error response with the specified status code, message, and error details.
func NewErrorResponse(ec echo.Context, statusCode int, message string, err error) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ec.Response().WriteHeader(statusCode)

	return ec.JSON(statusCode, Error{
		Message: message,
		Status:  statusCode,
		Error:   err,
	})
}

// NewSuccessResponse sends a JSON success response with data payload.
func NewSuccessResponse(ec echo.Context, statusCode int, data any) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ec.Response().WriteHeader(statusCode)

	return ec.JSON(statusCode, Success{
		Message: "OK",
		Data:    data,
	})
}

// NewSuccessMessageResponse sends a JSON success response with a success message.
func NewSuccessMessageResponse(ec echo.Context, statusCode int, message string) error {
	ec.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ec.Response().WriteHeader(statusCode)

	return ec.JSON(statusCode, Success{
		Message: message,
	})
}
