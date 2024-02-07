package adapters

import (
	"rest-project/src/config"
	"rest-project/src/domain/article"
)

//type Persistence struct {
//	UnitOfWork UnitOfWork
//}
//
//func NewPersistence(postgres *config.SPostgres) *Persistence {
//	return &Persistence{
//		UnitOfWork: newUnitOfWork(postgres),
//	}
//}

type sUnitOfWork struct {
	databaseContext   *config.SPostgres
	articleRepository ArticleRepository
}

func NewUnitOfWork(postgres *config.SPostgres) UnitOfWork {
	return &sUnitOfWork{
		databaseContext:   postgres,
		articleRepository: article.NewArticleRepository(postgres.DB),
	}
}

func (r *sUnitOfWork) ArticleRepository() ArticleRepository {
	return r.articleRepository
}

func (r *sUnitOfWork) Do(unitOfWorkBlock func(UnitOfWork) error) error {
	return r.databaseContext.Transaction(func(transaction *config.SPostgres) error {
		r.databaseContext = transaction
		return unitOfWorkBlock(r)
	})
}
