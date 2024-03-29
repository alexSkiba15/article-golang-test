package entities

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	Base
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (a Article) GetID() uuid.UUID {
	return a.ID
}

func NewArticle(id uuid.UUID, createdAt time.Time, updatedAt time.Time, title string, text string) *Article {
	return &Article{
		Base:  *NewBase(id, createdAt, updatedAt),
		Title: title,
		Text:  text,
	}
}
