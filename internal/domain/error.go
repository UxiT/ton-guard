package domain

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorCode int

var (
	InternalError    ErrorCode = 100
	EntityNotFound   ErrorCode = 101
	InvalidJWT       ErrorCode = 102
	ValueObjectError ErrorCode = 103

	ValidationFailed ErrorCode = 200

	InvalidUser        ErrorCode = 300
	InvalidCredentials ErrorCode = 301
	AccountExists      ErrorCode = 302
	InvalidProfileType ErrorCode = 303

	DomainOperationError       ErrorCode = 400
	ProviderAccountCreateError ErrorCode = 401
	SystemAccountCreateError   ErrorCode = 402
)

type ValidationConstraintError struct {
	Field string    `json:"field"`
	Code  ErrorCode `json:"code"`
}

func NewValidationConstraintError(field string) ValidationConstraintError {
	return ValidationConstraintError{
		Field: field,
		Code:  ValidationFailed,
	}
}

type ValidationError struct {
	Code   ErrorCode                   `json:"code"`
	Errors []ValidationConstraintError `json:"errors"`
}

func (e ValidationError) Error() string {
	var msg string

	for _, err := range e.Errors {
		msg += fmt.Sprintf("field: %s, code: %d,", err.Field, err.Code)
	}

	return msg
}

type ApplicationError struct {
	Err      error
	Code     ErrorCode
	HTTPCode int
}

func (e ApplicationError) Error() string {
	return e.Err.Error()
}

var (
	ErrInternal = ApplicationError{errors.New("internal error"), InternalError, http.StatusInternalServerError}

	ErrUnauthorized = ApplicationError{errors.New("unauthorized"), InvalidCredentials, http.StatusUnauthorized}
	ErrInvalidJWT   = ApplicationError{errors.New("invalid jwt token"), InvalidJWT, http.StatusUnauthorized}
	ErrInvalidUser  = ApplicationError{errors.New("invalid user"), InvalidUser, http.StatusUnauthorized}

	ErrCustomerNotFound = ApplicationError{errors.New("customer not found"), EntityNotFound, http.StatusNotFound}
	ErrAccountNotFound  = ApplicationError{errors.New("account not found"), EntityNotFound, http.StatusNotFound}

	ErrInvalidCredentials = ApplicationError{errors.New("invalid credentials"), InvalidCredentials, http.StatusUnauthorized}

	ErrInvalidProfileType = ApplicationError{errors.New("invalid account type"), InvalidProfileType, http.StatusUnauthorized}
	ErrAccountExists      = ApplicationError{errors.New("account already exists"), AccountExists, http.StatusUnprocessableEntity}
	ErrProviderAccount    = ApplicationError{errors.New("could not create account, please contact the support"), ProviderAccountCreateError, http.StatusInternalServerError}
	ErrSystemAccount      = ApplicationError{errors.New("could not create account, please contact the support"), SystemAccountCreateError, http.StatusInternalServerError}
)
