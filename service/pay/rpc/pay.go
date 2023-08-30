package main

import (
	"flag"
	"fmt"

	"github.com/fly602/mall/service/pay/rpc/internal/config"
	"github.com/fly602/mall/service/pay/rpc/internal/server"
	"github.com/fly602/mall/service/pay/rpc/internal/svc"
	"github.com/fly602/mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/pay.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pay.RegisterPayServer(grpcServer, server.NewPayServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
