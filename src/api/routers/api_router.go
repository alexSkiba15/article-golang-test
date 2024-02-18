package routers

import (
	"github.com/gin-gonic/gin"
	"rest-project/src/config"
	"time"
)

func SetupRouter(conf *config.Config, db *config.SPostgres, timeout time.Duration, engine *gin.Engine) {
	api := engine.Group("/api")
	NewArticleRouter(conf, db, timeout, api)
}
