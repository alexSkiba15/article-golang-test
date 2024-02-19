package main

import (
	"rest-project/src/api/routers"
	"rest-project/src/config"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, _ := config.NewConfig()
	db := config.DBConnection(conf.PostgresConfig)
	defer db.Close()
	timeout := time.Duration(60) * time.Second
	engine := gin.Default()
	engine.Use(gin.Logger())
	routers.SetupRouter(conf, db, timeout, engine)

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
