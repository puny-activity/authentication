package account

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID        *uuid.UUID
	Username  string
	Nickname  string
	CreatedAt carbon.Carbon
}

type Account struct {
	User
	Roles      []role.Role
	LastActive carbon.Carbon
}

type ToCreate struct {
	User
	Password string
}

type ToCreateWithHashedPassword struct {
	User
	HashedPassword string
}

func (e *User) Validate() error {
	if len(e.Username) == 0 {
		return errs.NotProvidedUsername
	}

	if len(e.Nickname) == 0 {
		return errs.NotProvidedNickname
	}

	if strings.ContainsAny(e.Username, " \\/") {
		return errs.InvalidUsername
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
	hashedPassword := string(hashedBytes)

	return ToCreateWithHashedPassword{
		User:           e.User,
		HashedPassword: hashedPassword,
	}, nil
}
