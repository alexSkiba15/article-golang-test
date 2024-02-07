package article

import (
	"rest-project/src/adapters"
	"rest-project/src/adapters/models"
	"rest-project/src/domain/entities"

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
