package handler

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc/client/pool"
	"grpc/service/agent"
)

const (
	defaultPoolSize = 20
	testTarget      = "127.0.0.1:13004"
)

func makeServiceClient(cc *grpc.ClientConn) interface{} {
	return agent.NewAgentServiceClient(cc)
}

type Handler struct {
	p   *pool.Pool
	ctx context.Context
}

func New(ctx context.Context) *Handler {
	return &Handler{
		p:   pool.New(ctx, defaultPoolSize, testTarget, makeServiceClient),
		ctx: ctx,
	}
}
