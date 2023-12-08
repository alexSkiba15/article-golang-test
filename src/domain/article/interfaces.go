package article

import (
	"github.com/google/uuid"
	"rest-project/src/domain/entities"
)

type Input struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type UseCases interface {
	GetAllArticleData() (*[]entities.Article, error)
	Create(i Input) (*entities.Article, error)
	Delete(id uuid.UUID) error
	GetArticle(articleId uuid.UUID) (*entities.Article, error)
	UpdateArticle(articleId uuid.UUID, i Input) (*entities.Article, error)
}
