package interr

import (
	"errors"
	"github.com/puny-activity/authentication/pkg/loclzr"
	"net/http"
)

type internalError struct {
	Error          error
	InternalCode   string
	HTTPStatusCode int
}

type ErrorStorage struct {
	localizer *loclzr.Localizer
}

func NewErrorStorage(localizer *loclzr.Localizer) *ErrorStorage {
	return &ErrorStorage{
		localizer: localizer,
	}
}

var (
	EmptyAccessToken = errors.New("empty access token")

	FailedToDecodeRequestBody = errors.New("failed to decode request body")
	NotProvidedUsername       = errors.New("not provided username")
	NotProvidedPassword       = errors.New("not provided password")
	InvalidUsername           = errors.New("invalid username")
	InvalidPassword           = errors.New("invalid password")

	UsernameAlreadyTaken = errors.New("username already taken")

	DatabaseFailedToExecuteQuery = errors.New("failed to execute query")

	Unexpected = errors.New("unexpected error")
)

var (
	errEmptyAccessToken = internalError{
		Error:          EmptyAccessToken,
		InternalCode:   "ATH-1",
		HTTPStatusCode: http.StatusUnauthorized,
	}

	errFailedToDecodeRequestBody = internalError{
		Error:          FailedToDecodeRequestBody,
		InternalCode:   "RDT-1",
		HTTPStatusCode: http.StatusBadRequest,
	}

	errNotProvidedUsername = internalError{
		Error:          NotProvidedUsername,
		InternalCode:   "RDT-2",
		HTTPStatusCode: http.StatusBadRequest,
	}
	errNotProvidedPassword = internalError{
		Error:          NotProvidedPassword,
		InternalCode:   "RDT-3",
		HTTPStatusCode: http.StatusBadRequest,
	}

	errInvalidUsername = internalError{
		Error:          InvalidUsername,
		InternalCode:   "RDT-4",
		HTTPStatusCode: http.StatusBadRequest,
	}
	errInvalidPassword = internalError{
		Error:          InvalidPassword,
		InternalCode:   "RDT-5",
		HTTPStatusCode: http.StatusBadRequest,
	}

	errUsernameAlreadyTaken = internalError{
		Error:          UsernameAlreadyTaken,
		InternalCode:   "CFL-1",
		HTTPStatusCode: http.StatusBadRequest,
	}

	errDatabaseFailedToExecuteQuery = internalError{
		Error:          DatabaseFailedToExecuteQuery,
		InternalCode:   "DTB-1",
		HTTPStatusCode: http.StatusInternalServerError,
	}

	errUnexpectedError = internalError{
		Error:          Unexpected,
		InternalCode:   "SRV-1",
		HTTPStatusCode: http.StatusInternalServerError,
	}
)

func (e *ErrorStorage) getInternalError(err error) internalError {
	switch {
	case errors.Is(err, errEmptyAccessToken.Error):
		return errEmptyAccessToken

	case errors.Is(err, errFailedToDecodeRequestBody.Error):
		return errFailedToDecodeRequestBody

	case errors.Is(err, errNotProvidedUsername.Error):
		return errNotProvidedUsername
	case errors.Is(err, errNotProvidedPassword.Error):
		return errNotProvidedPassword

	case errors.Is(err, errInvalidUsername.Error):
		return errInvalidUsername
	case errors.Is(err, errInvalidPassword.Error):
		return errInvalidPassword

	case errors.Is(err, errUsernameAlreadyTaken.Error):
		return errUsernameAlreadyTaken

	case errors.Is(err, errDatabaseFailedToExecuteQuery.Error):
		return errDatabaseFailedToExecuteQuery

	case errors.Is(err, errUnexpectedError.Error):
		return errUnexpectedError

	default:
		return errUnexpectedError
	}
}

func (e *ErrorStorage) GetHTTPStatusCode(err error) int {
	return e.getInternalError(err).HTTPStatusCode
}

func (e *ErrorStorage) GetInternalCode(err error) string {
	return e.getInternalError(err).InternalCode
}

func (e *ErrorStorage) GetMessage(lang string, err error) string {
	internalCode := e.getInternalError(err).InternalCode

	message, err := e.localizer.Localize(lang, internalCode)
	if err != nil {
		return e.localizer.English(internalCode)
	}

	return message
}
