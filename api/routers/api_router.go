package routers

import (
	"time"

	"github.com/alexSkiba15/article-golang-test/config"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(conf *config.Config, db *config.SPostgres, timeout time.Duration, app *fiber.App) {
	api := app.Group("/api")
	NewArticleRouter(conf, db, timeout, api)
}
