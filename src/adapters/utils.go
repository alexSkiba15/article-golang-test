package adapters

import (
	"github.com/google/uuid"
)

func UUIDGenerator() uuid.UUID {
	generatedUuid, _ := uuid.NewUUID()
	return generatedUuid
}
