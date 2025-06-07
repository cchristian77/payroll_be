package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var v *Validator

type Validator struct {
	validate *validator.Validate
}

func RegisterValidator() *Validator {
	if v == nil {
		v = &Validator{
			validate: validator.New(),
		}
	}

	return v
}

func (v *Validator) Validate(request any) error {
	if err := v.validate.Struct(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
