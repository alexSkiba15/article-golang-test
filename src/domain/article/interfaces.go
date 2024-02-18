package article

import (
	"context"
	"github.com/google/uuid"
	"rest-project/src/domain/entities"
)

type Input struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type UseCases interface {
	GetAllArticleData(ctx context.Context) (*[]entities.Article, error)
	Create(ctx context.Context, i Input) (*entities.Article, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetArticle(ctx context.Context, articleId uuid.UUID) (*entities.Article, error)
	UpdateArticle(ctx context.Context, articleId uuid.UUID, i Input) (*entities.Article, error)
}
