package loginattempt

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/pkg/util"
)

type ID uuid.UUID

type Row struct {
	ID          *ID
	Email       email.Email
	Success     bool
	AttemptedAt carbon.Carbon
}

func (e Row) GenerateID() Row {
	e.ID = util.ToPointer(ID(uuid.New()))
	return e
}
