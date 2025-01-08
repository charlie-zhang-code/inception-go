package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	Consul consul.Conf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	MySQL struct {
		DataSource   string
		MaxIdleConns int `json:",default=10"`
		MaxOpenConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}

	Cache redis.RedisConf

	Default struct {
		Password string
	}

	MemberRpc zrpc.RpcClientConf
}
