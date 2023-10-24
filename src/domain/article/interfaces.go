package article

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
}

func (a Article) GetId() uuid.UUID {
	return a.ID
}

type Input struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type UseCases interface {
	GetAllArticleData() (*[]Article, error)
	Create(i Input) (*Article, error)
	Delete(id uuid.UUID) error
	GetArticle(articleId uuid.UUID) (*Article, error)
	UpdateArticle(articleId uuid.UUID, i Input) (*Article, error)
}
