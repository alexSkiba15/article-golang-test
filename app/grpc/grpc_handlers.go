package grpcapp

import (
	"context"
	"fmt"
	timergrpc "rest-project/api/grpc/gen"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type serverGRPCAPI struct {
	timergrpc.UnimplementedAPITimerServer
}

func Register(gRPC *grpc.Server) {
	timergrpc.RegisterAPITimerServer(gRPC, &serverGRPCAPI{})
}

func (s *serverGRPCAPI) CallbackConfig(
	ctx context.Context, req *timergrpc.CallbackConfigRequest,
) (*timergrpc.CallbackConfigResponse, error) {
	fmt.Println("CallbackConfig")
	idCallback, _ := uuid.NewUUID()
	return &timergrpc.CallbackConfigResponse{
		Id:     idCallback.String(),
		Url:    *req.Url,
		Params: req.Params,
		Type:   req.Type,
	}, nil
}

func (s *serverGRPCAPI) Timer(ctx context.Context, req *timergrpc.TimerRequest) (*timergrpc.TimerResponse, error) {
	panic("implement")
}
