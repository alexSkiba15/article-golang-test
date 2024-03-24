package main

import (
	"strconv"

	"github.com/alexSkiba15/article-golang-test/app"
	"github.com/alexSkiba15/article-golang-test/config"
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
