package entities

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewBase(ID uuid.UUID, createdAt time.Time, updatedAt time.Time) *Base {
	return &Base{
		ID:        ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
