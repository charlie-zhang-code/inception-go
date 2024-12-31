package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"member/proto/model"
	"member/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	SysUserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		SysUserModel: model.NewSysUserModel(conn),
	}
}
