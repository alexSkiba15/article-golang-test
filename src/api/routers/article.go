package routers

import (
	"rest-project/src/api/controllers"
	"rest-project/src/config"
	"rest-project/src/domain/article"
	"time"

	"github.com/gin-gonic/gin"
)

func NewArticleRouter(conf *config.Config, db *config.SPostgres, timeout time.Duration, group *gin.RouterGroup) {
	repository := article.NewArticleRepository(db.DB)
	articleUseCases := article.NewArticleUseCasesImpl(*repository)
	articleController := controllers.NewArticleHandler(articleUseCases)

	articleGroup := group.Group("/article")
	{
		articleGroup.GET("", articleController.GetAll)
		articleGroup.POST("", articleController.Create)
		articleGroup.DELETE("/:articleID", articleController.Delete)
		articleGroup.GET("/:articleID", articleController.GetArticleById)
		articleGroup.PUT("/:articleID", articleController.UpdateArticle)
	}
}
