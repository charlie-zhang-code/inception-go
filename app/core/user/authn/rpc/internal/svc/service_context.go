package svc

import (
	"authn/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"member/rpc/member"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	DB            *gorm.DB
	Cache         *redis.Redis
	MemberService member.Member
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.MySQL.DataSource))
	sdb, _ := db.DB()
	sdb.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	sdb.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	sdb.SetConnMaxLifetime(time.Second * time.Duration(c.MySQL.MaxLifetime))

	cache := redis.MustNewRedis(redis.RedisConf{
		Host:        c.Cache.Host,
		Type:        c.Cache.Type,
		Pass:        c.Cache.Pass,
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})

	return &ServiceContext{
		Config:        c,
		DB:            db,
		Cache:         cache,
		MemberService: member.NewMember(zrpc.MustNewClient(c.MemberRpc)),
	}
}
