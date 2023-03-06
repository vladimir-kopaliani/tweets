package apperrors

import "errors"

var (
	// internal errors
	ErrLoggerIsNotSet     = errors.New("logger is not set")
	ErrAddressIsNotSet    = errors.New("address is not set")
	ErrServiceIsNotSet    = errors.New("service is not set")
	ErrRepositoryIsNotSet = errors.New("repository is not set")
	ErrExtractUserContext = errors.New("can not extract user context")
	ErrJWTEmpty           = errors.New("JWT token is empty")

	// public errors
	ErrInternal               = errors.New("internal error")
	ErrWrongCredentials       = errors.New("wrong credentials")
	ErrJWTInvalid             = errors.New("JWT token is not valid")
	ErrAuthenticationRequired = errors.New("authentication is required")
)

func IsPublicError(err error) bool {
	switch err {
	case ErrInternal, ErrWrongCredentials,
		ErrJWTInvalid, ErrAuthenticationRequired,
		ErrJWTEmpty:
		return true
	}

	return false
}
