package article

import (
	"gorm.io/gorm"
	"rest-project/src/adapters"
	"rest-project/src/domain/entities"
)

type RepositoryArticle interface {
	*adapters.Repository[adapters.Article, entities.Article]
}

type ArticleRepository struct {
	*adapters.Repository[adapters.Article, entities.Article]
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{adapters.NewRepository[adapters.Article, entities.Article](db)}
}
