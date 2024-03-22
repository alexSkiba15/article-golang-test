package main

import (
	"rest-project/api/routers"
	"rest-project/config"
	"time"

	"github.com/bytedance/sonic"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	conf := config.NewConfig()
	db := config.DBConnection(conf.PostgresDB)
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
