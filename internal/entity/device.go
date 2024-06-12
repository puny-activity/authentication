package entity

import "github.com/google/uuid"

type Device struct {
	ID          *string
	Fingerprint string
}

func (e *Device) GenerateID() {
	generatedID := uuid.New().String()
	e.ID = &generatedID
}
