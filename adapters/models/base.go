package models

import (
	"rest-project/domain/entities"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b BaseModel) ToEntity() *entities.Base {
	return entities.NewBase(
		b.ID,
		b.CreatedAt,
		b.UpdatedAt,
	)
}

func (b BaseModel) FromEntity(entity *entities.Base) any {
	b.ID = entity.ID
	b.UpdatedAt = entity.UpdatedAt
	b.CreatedAt = entity.CreatedAt

	return b
}
