package device

import (
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/pkg/util"
)

type ID uuid.UUID

type Device struct {
	ID          *ID
	Name        string
	Fingerprint string
}

func (e Device) GenerateID() Device {
	e.ID = util.ToPointer(ID(uuid.New()))
	return e
}
