package entities

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewBase(id uuid.UUID, createdAt time.Time, updatedAt time.Time) *Base {
	return &Base{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
