package models

import (
	"fmt"
	"github.com/google/uuid"
	"rest-project/src/domain/entities"
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
		fmt.Println(err)
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

func (a Article) FromEntity(entity *entities.Article) any {
	a.ID = entity.ID
	a.Title = entity.Title
	a.Text = entity.Text
	a.BaseModel.FromEntity(&entity.Base)
	return a
}
