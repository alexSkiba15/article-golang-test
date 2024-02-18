package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func NewGRPCApp(port int) *App {
	gRPCServer := grpc.NewServer()
	Register(gRPCServer)
	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Start() error {
	const op = "app.App.Start"
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: error in listen %w", op, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: error in serve %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
