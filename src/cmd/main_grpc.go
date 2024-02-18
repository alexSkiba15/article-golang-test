package main

import (
	"rest-project/src/app"
	"rest-project/src/config"
	"strconv"
)

func main() {
	conf, _ := config.NewConfig()
	port, _ := strconv.Atoi(conf.GRPCServerConfig.Port)

	application := app.New(port)
	err := application.GRPCServer.Start()
	if err != nil {
		panic(err)
		return
	}
}
