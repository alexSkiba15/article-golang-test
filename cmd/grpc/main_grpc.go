package main

import (
	"rest-project/app"
	"rest-project/config"
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
