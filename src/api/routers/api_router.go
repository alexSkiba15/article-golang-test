package routers

import (
	"rest-project/src/config"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(conf *config.Config, db *config.SPostgres, timeout time.Duration, engine *gin.Engine) {
	api := engine.Group("/api")
	NewArticleRouter(conf, db, timeout, api)
}
