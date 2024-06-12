package entity

import "github.com/golang-module/carbon"

type Account struct {
	ID       string
	Username string
	Roles    []Role
}

type AccountCreateRequest struct {
	Account
	Password string
}

type AccountCreateRequestWithHashedPassword struct {
	Account
	HashedPassword string
}

type PublicAccount struct {
	Account
	CreatedAt  carbon.Carbon
	LastActive carbon.Carbon
}
