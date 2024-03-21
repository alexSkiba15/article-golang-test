package routers

import (
	"rest-project/src/api/controllers"
	"rest-project/src/config"
	"rest-project/src/domain/article"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewArticleRouter(conf *config.Config, db *config.SPostgres, timeout time.Duration, group fiber.Router) {
	repository := article.NewArticleRepository(db.DB)
	articleUseCases := article.NewArticleUseCasesImpl(*repository)
	articleController := controllers.NewArticleHandler(articleUseCases)

	articleGroup := group.Group("/article")
	{
		articleGroup.Get("", articleController.GetAll)
		articleGroup.Post("", articleController.Create)
		articleGroup.Delete("/:articleID", articleController.Delete)
		articleGroup.Get("/:articleID", articleController.GetArticleById)
		articleGroup.Put("/:articleID", articleController.UpdateArticle)
	}
}
