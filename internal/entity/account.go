package entity

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/interr"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Account struct {
	ID         *string
	Username   string
	Roles      []Role
	CreatedAt  carbon.Carbon
	LastActive carbon.Carbon
}

type AccountCreateRequest struct {
	Account
	Password string
}

type AccountCreateRequestWithHashedPassword struct {
	Account
	HashedPassword string
}

func (e *Account) GenerateID() {
	generatedID := uuid.New().String()
	e.ID = &generatedID
}

func (e *Account) Validate() error {
	if len(e.Username) == 0 {
		return interr.NotProvidedUsername
	}

	if strings.ContainsAny(e.Username, " \\/") {
		return interr.InvalidUsername
	}

	return nil
}

func (e *AccountCreateRequest) Validate() error {
	err := e.Account.Validate()
	if err != nil {
		return err
	}

	if len(e.Password) == 0 {
		return interr.NotProvidedPassword
	}

	if strings.ContainsAny(e.Password, " \\/") || len(e.Password) < 8 {
		return interr.InvalidPassword
	}

	return nil
}

func (e *AccountCreateRequest) HashPassword() (AccountCreateRequestWithHashedPassword, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return AccountCreateRequestWithHashedPassword{}, err
	}
	hashedPassword := string(hashedBytes)

	return AccountCreateRequestWithHashedPassword{
		Account:        e.Account,
		HashedPassword: hashedPassword,
	}, nil
}
