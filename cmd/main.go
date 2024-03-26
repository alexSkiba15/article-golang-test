package main

import (
	"time"

	"github.com/alexSkiba15/article-golang-test/api/routers"
	"github.com/alexSkiba15/article-golang-test/config"

	"github.com/bytedance/sonic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	conf := config.New()
	db := config.NewDBConnection(conf.PostgresDB)
	defer db.Close()
	timeout := time.Duration(60) * time.Second
	engine := fiber.New(
		fiber.Config{JSONDecoder: sonic.Unmarshal, JSONEncoder: sonic.Marshal},
	)
	engine.Use(cors.New())
	routers.SetupRouter(conf, db, timeout, engine)
	err := engine.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
