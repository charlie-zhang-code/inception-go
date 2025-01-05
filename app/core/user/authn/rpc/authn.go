package main

import (
	"flag"
	"fmt"

	"authn/rpc/internal/config"
	jwttokenServer "authn/rpc/internal/server/jwttoken"
	opaquetokenServer "authn/rpc/internal/server/opaquetoken"
	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/authn.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterJwtTokenServer(grpcServer, jwttokenServer.NewJwtTokenServer(ctx))
		pb.RegisterOpaqueTokenServer(grpcServer, opaquetokenServer.NewOpaqueTokenServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
