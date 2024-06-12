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
	Unauthorized = errors.New("unauthorized")

	InternalServer = errors.New("internal server error")
)

var (
	ErrEmptyAccessToken = internalError{
		Error:          Unauthorized,
		InternalCode:   "ATH-001",
		HTTPStatusCode: http.StatusUnauthorized,
	}
	ErrUnexpectedError = internalError{
		Error:          InternalServer,
		InternalCode:   "SRV-001",
		HTTPStatusCode: http.StatusUnauthorized,
	}
)

func (e *ErrorStorage) getInternalError(err error) internalError {
	switch {
	case errors.Is(err, ErrEmptyAccessToken.Error):
		return ErrEmptyAccessToken
	default:
		return ErrUnexpectedError
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
