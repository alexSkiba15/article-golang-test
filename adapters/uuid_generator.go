package adapters

import (
	"github.com/google/uuid"
)

func UUIDGenerator() uuid.UUID {
	generatedUUID, _ := uuid.NewUUID()
	return generatedUUID
}
