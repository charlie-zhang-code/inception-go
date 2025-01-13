package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Nacos struct {
		Host                string
		Port                uint64
		Namespace           string
		NotLoadCacheAtStart bool
		LogLevel            string
		ServiceName         string
	}
}
