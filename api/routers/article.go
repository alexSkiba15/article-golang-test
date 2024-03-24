package routers

import (
	"time"

	"github.com/alexSkiba15/article-golang-test/api/controllers"
	"github.com/alexSkiba15/article-golang-test/config"
	"github.com/alexSkiba15/article-golang-test/domain/article"

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
