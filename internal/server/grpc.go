package server

import (
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "lottery/api/lottery/v1"
	"lottery/internal/conf"
	"lottery/internal/service"
)

// NewGRPCServer
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: NewGRPCServer new a gRPC server.
//	@param c
//	@param greeter
//	@return *grpc.Server
func NewGRPCServer(c *conf.Server, greeter *service.LotteryService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			mmd.Server(),
			MiddlewareTraceID(),
			MiddlewareLog(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterLotteryServer(srv, greeter)
	return srv
}
