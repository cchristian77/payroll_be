package middleware

import (
	"context"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// RequestID adds a middleware to ensure that every request has a unique request ID, using `X-Request-ID` header or generating one.
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			ctx := ec.Request().Context()

			requestID := ec.Request().Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			ctx = context.WithValue(ctx, enums.RequestIDCtxKey, requestID)
			ec.SetRequest(ec.Request().WithContext(ctx))

			return next(ec)
		}
	}
}
