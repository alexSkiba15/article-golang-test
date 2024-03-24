package article

import (
	"github.com/alexSkiba15/article-golang-test/adapters"
	"github.com/alexSkiba15/article-golang-test/adapters/models"
	"github.com/alexSkiba15/article-golang-test/domain/entities"

	"gorm.io/gorm"
)

type RepositoryArticle interface {
	*adapters.Repository[models.Article, entities.Article]
}

type Repository struct {
	*adapters.Repository[models.Article, entities.Article]
}

func NewArticleRepository(db *gorm.DB) *Repository {
	return &Repository{adapters.NewRepository[models.Article, entities.Article](db)}
}
