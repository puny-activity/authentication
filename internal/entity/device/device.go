package device

import "github.com/google/uuid"

type Device struct {
	ID          *uuid.UUID
	Fingerprint string
}
