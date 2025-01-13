package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"

	"member/rpc/internal/config"
	"member/rpc/internal/server"
	"member/rpc/internal/svc"
	"member/rpc/pb"

	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/member.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMemberServer(grpcServer, server.NewMemberServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// register service to nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Nacos.Host, c.Nacos.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         c.Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: c.Nacos.NotLoadCacheAtStart,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            c.Nacos.LogLevel,
	}

	opts := nacos.NewNacosConfig(c.Nacos.ServiceName, c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
