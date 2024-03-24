package app

import (
	grpcapp "github.com/alexSkiba15/article-golang-test/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	grpcPort int,
) *App {
	grpcApp := grpcapp.NewGRPCApp(grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
