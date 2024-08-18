package password

import (
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Password struct {
	password string
}

func New(password string) (Password, error) {
	if password == "" {
		return Password{}, errs.InvalidPassword
	}

	if len([]rune(password)) < 4 {
		return Password{}, errs.InvalidPassword
	}

	if strings.ContainsAny(password, " \\/") {
		return Password{}, errs.InvalidPassword
	}

	return Password{
		password: password,
	}, nil
}

type Hashed struct {
	hashedPassword string
}

func NewHashed(hashedPassword string) Hashed {
	return Hashed{
		hashedPassword: hashedPassword,
	}
}

func (e *Hashed) String() string {
	return e.hashedPassword
}

func (e *Password) Hash() (Hashed, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(e.password), bcrypt.DefaultCost)
	if err != nil {
		return Hashed{}, werr.WrapSE("failed to generate hashed password", err)
	}
	passwordHash := string(hashedBytes)
	return NewHashed(passwordHash), nil
}

func IsMatch(password Password, hashedPassword Hashed) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.hashedPassword), []byte(password.password))
	return err == nil
}
