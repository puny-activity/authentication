package account

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/pkg/util"
)

type ID uuid.UUID

type Account struct {
	ID         *ID
	Email      email.Email
	Nickname   string
	Role       role.Role
	CreatedAt  carbon.Carbon
	LastActive *carbon.Carbon
}

func (e Account) GenerateID() Account {
	e.ID = util.ToPointer(ID(uuid.New()))
	return e
}
