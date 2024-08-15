package account

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"strings"
)

type User struct {
	ID        *uuid.UUID
	Email     string
	Nickname  string
	Role      role.Role
	CreatedAt carbon.Carbon
}

type Account struct {
	User
	LastActive carbon.Carbon
}

type ToCreate struct {
	User
	Password string
}

type ToCreateWithHashedPassword struct {
	User
	PasswordHash string
}

func (e *User) Validate() error {
	if len(e.Email) == 0 {
		return errs.NotProvidedEmail
	}

	if len(e.Nickname) == 0 {
		return errs.NotProvidedNickname
	}

	_, err := mail.ParseAddress(e.Email)
	if err != nil {
		return werr.WrapES(errs.InvalidEmail, err.Error())
	}

	return nil
}

func (e *Account) Validate() error {
	return e.User.Validate()
}

func (e *ToCreate) Validate() error {
	err := e.User.Validate()
	if err != nil {
		return err
	}

	if strings.ContainsAny(e.Password, " \\/") || len(e.Password) < 8 {
		return errs.InvalidPassword
	}

	return nil
}

func (e *ToCreate) HashPassword() (ToCreateWithHashedPassword, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return ToCreateWithHashedPassword{}, werr.WrapSE("failed to generate hashed password", err)
	}
	passwordHash := string(hashedBytes)

	return ToCreateWithHashedPassword{
		User:         e.User,
		PasswordHash: passwordHash,
	}, nil
}
