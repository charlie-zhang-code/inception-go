package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"member/proto/model"
	"member/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Member model.SysMemberModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Member: model.NewSysMemberModel(sqlx.NewMysql(c.DataSource)),
		//Member: model.NewSysMemberModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
