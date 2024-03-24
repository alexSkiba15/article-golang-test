package models

import (
	"fmt"

	"github.com/alexSkiba15/article-golang-test/domain/entities"

	"github.com/google/uuid"
)

type Article struct {
	Title string `gorm:"type:varchar(256);not null;"`
	Text  string `gorm:"type:varchar(2048);not null;"`
	BaseModel
}

func (a Article) GetID() uuid.UUID {
	return a.ID
}

func (a Article) GenUUID() (uuid.UUID, error) {
	generatedUUID, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("error generating UUID: %v", err)
	}

	return generatedUUID, nil
}

func (a Article) ToEntity() *entities.Article {
	return entities.NewArticle(
		a.ID,
		a.CreatedAt,
		a.UpdatedAt,
		a.Title,
		a.Text,
	)
}

func (a Article) FromEntity(entity *entities.Article) Article {
	a.ID = entity.ID
	a.Title = entity.Title
	a.Text = entity.Text
	a.BaseModel.FromEntity(&entity.Base)

	return a
}
