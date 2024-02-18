package app

import (
	grpcapp "rest-project/src/app/grpc"
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
