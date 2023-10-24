package article

import (
	"gorm.io/gorm"
	"rest-project/src/adapters"
)

type RepositoryArticle interface {
	*adapters.Repository[adapters.Article]
}

type ArticleRepository struct {
	*adapters.Repository[adapters.Article]
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{adapters.NewRepository[adapters.Article](db)}
}
