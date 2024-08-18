package email

import (
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
	"net/mail"
)

type Email struct {
	email string
}

func New(email string) (Email, error) {
	if email == "" {
		return Email{}, errs.NotProvidedEmail
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return Email{}, werr.WrapES(errs.InvalidEmail, err.Error())
	}

	return Email{
		email: email,
	}, nil
}

func (e Email) String() string {
	return e.email
}
