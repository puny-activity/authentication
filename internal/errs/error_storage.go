package errs

import (
	"errors"
	"github.com/puny-activity/authentication/pkg/loclzr"
)

type internalError struct {
	error            error
	localizationCode string
	httpStatusCode   int
}

type Storage struct {
	localizer *loclzr.Localizer
}

func NewStorage(localizer *loclzr.Localizer) *Storage {
	return &Storage{
		localizer: localizer,
	}
}

func (s *Storage) getInternalError(err error) internalError {
	for i := range errorList {
		if errors.Is(err, errorList[i].error) {
			return errorList[i]
		}
	}
	return unexpectedError
}

func (s *Storage) HTTPStatusCode(err error) int {
	return s.getInternalError(err).httpStatusCode
}

func (s *Storage) LocalizationCode(err error) string {
	return s.getInternalError(err).localizationCode
}

func (s *Storage) Message(lang string, err error) string {
	return s.localizer.Localize(lang, s.getInternalError(err).localizationCode)
}
