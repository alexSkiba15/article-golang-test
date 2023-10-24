package routers

import (
	"github.com/gin-gonic/gin"
	"rest-project/src/api/controllers"
)

func SetupRouter(articleController controllers.ArticleHandler) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	api := router.Group("/api")
	{
		article := api.Group("/article")
		article.GET("", articleController.GetAll)
		article.POST("", articleController.Create)
		article.DELETE("/:articleID", articleController.Delete)
		article.GET("/:articleID", articleController.GetArticleById)
		article.PUT("/:articleID", articleController.UpdateArticle)
	}

	return router
}
