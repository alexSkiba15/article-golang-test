package adapters

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ArticlePK struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	Title string `gorm:"type:varchar(256);not null;"`
	Text  string `gorm:"type:varchar(2048);not null;"`
	BaseModel
	ArticlePK
}

func (a Article) GetId() uuid.UUID {
	return a.ArticlePK.ID
}

func (a Article) GenUUID() (uuid.UUID, error) {
	generatedUuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}
	return generatedUuid, nil
}
