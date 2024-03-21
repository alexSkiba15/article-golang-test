package main

import (
	"rest-project/src/app"
	"rest-project/src/config"
	"strconv"
)

func main() {
	conf := config.NewConfig()
	port, _ := strconv.Atoi(conf.GRPCServer.Port)

	application := app.New(port)
	err := application.GRPCServer.Start()
	if err != nil {
		panic(err)
	}
}
