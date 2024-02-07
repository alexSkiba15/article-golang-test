package main

import (
	"context"
	"rest-project/src/adapters"
	"rest-project/src/api/controllers"
	"rest-project/src/api/routers"
	"rest-project/src/config"
	"rest-project/src/domain/article"
)

func main() {
	conf, _ := config.NewConfig()
	db := config.DBConnection(conf.PostgresConfig)
	uow := adapters.NewUnitOfWork(db)

	//test := article.NewArticleRepository(uow)
	ctx := context.Background()
	articleUseCases := article.NewArticleUseCasesImpl(*test, ctx)
	articleController := controllers.NewArticleHandler(articleUseCases)
	r := routers.SetupRouter(articleController)
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
