package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Cache redis.RedisConf

	MySQL struct {
		DataSource   string
		MaxIdleConns int `json:",default=10"`
		MaxOpenConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Consul consul.Conf
}
