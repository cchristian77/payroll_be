package helper

import (
	"errors"
)

// HTTP ERROR
var (
	// InternalServerErr will throw if any the Internal Server Error happen
	InternalServerErr = errors.New("Internal Server Error")

	// NotFoundErr will throw if the requested item is not exists
	NotFoundErr = errors.New("Requested data is not found")

	// ConflictErr will throw if the current action already exists
	ConflictErr = errors.New("Requested data already exist")

	// BadParamInputErr will throw if the given request-body or params is not valid
	BadParamInputErr = errors.New("Requested parameters are not valid")

	// ForbiddenErr will throw if the current request is forbidden
	ForbiddenErr = errors.New("Forbidden Access")

	// UnauthorizedErr will throw if the current request is unauthorized
	UnauthorizedErr = errors.New("Unauthorized")

	// IncorrectCredentialErr will throw if the email or password credential is incorrect
	IncorrectCredentialErr = errors.New("Login failed. Email or password is incorrect.")
)

// AUTH ERROR
var (
	InvalidTokenErr = errors.New("invalid token")
	ExpiredTokenErr = errors.New("expired token")
)

type BusinessValidationErr struct {
	Message string
}

func (e BusinessValidationErr) Error() string {
	return e.Message
}

func NewBusinessValidationErr(message string) error {
	return BusinessValidationErr{
		Message: message,
	}
}
