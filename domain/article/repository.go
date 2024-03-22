package article

import (
	"rest-project/adapters"
	"rest-project/adapters/models"
	"rest-project/domain/entities"

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
