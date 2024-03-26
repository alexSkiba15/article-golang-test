//go:build wireinject
// +build wireinject

package di

import (
	"github.com/alexSkiba15/article-golang-test/adapters"
	"github.com/alexSkiba15/article-golang-test/config"
	"github.com/alexSkiba15/article-golang-test/domain/article"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var providerSet wire.ProviderSet = wire.NewSet(
	config.NewGormDB,
	config.NewDBConnection,
	config.NewPostgresDB,
	article.NewArticleRepository,
	article.NewArticleUseCasesImpl,
	wire.Bind(new(adapters.ArticleRepository), new(*article.Repository)),
)

func InitializeConfig() *config.Config {
	wire.Build(config.New)
	return &config.Config{}
}

func InitializeDBConnection() *config.SPostgres {
	wire.Build(providerSet)
	return &config.SPostgres{}
}

func InitializeGormDB() *gorm.DB {
	wire.Build(providerSet)
	return &gorm.DB{}
}

func InitializeArticleRepository() *article.Repository {
	wire.Build(providerSet)
	return &article.Repository{}
}

func InitializeArticleUseCase() *article.UseCasesImpl {
	wire.Build(providerSet)
	return &article.UseCasesImpl{}
}
