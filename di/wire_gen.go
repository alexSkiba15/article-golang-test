// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/alexSkiba15/article-golang-test/adapters"
	"github.com/alexSkiba15/article-golang-test/config"
	"github.com/alexSkiba15/article-golang-test/domain/article"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeConfig() *config.Config {
	configConfig := config.New()
	return configConfig
}

func InitializeDBConnection() *config.SPostgres {
	postgresDB := config.NewPostgresDB()
	sPostgres := config.NewDBConnection(postgresDB)
	return sPostgres
}

func InitializeGormDB() *gorm.DB {
	postgresDB := config.NewPostgresDB()
	sPostgres := config.NewDBConnection(postgresDB)
	db := config.NewGormDB(sPostgres)
	return db
}

func InitializeArticleRepository() *article.Repository {
	postgresDB := config.NewPostgresDB()
	sPostgres := config.NewDBConnection(postgresDB)
	db := config.NewGormDB(sPostgres)
	repository := article.NewArticleRepository(db)
	return repository
}

func InitializeArticleUseCase() *article.UseCasesImpl {
	postgresDB := config.NewPostgresDB()
	sPostgres := config.NewDBConnection(postgresDB)
	db := config.NewGormDB(sPostgres)
	repository := article.NewArticleRepository(db)
	useCasesImpl := article.NewArticleUseCasesImpl(repository)
	return useCasesImpl
}

// wire.go:

var providerSet wire.ProviderSet = wire.NewSet(config.NewGormDB, config.NewDBConnection, config.NewPostgresDB, article.NewArticleRepository, article.NewArticleUseCasesImpl, wire.Bind(new(adapters.ArticleRepository), new(*article.Repository)))