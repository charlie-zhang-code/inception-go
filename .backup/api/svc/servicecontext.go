package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"member/api/internal/config"
	"member/rpc/client/sysmemberservice"
)

type ServiceContext struct {
	Config       config.Config
	MemberClient sysmemberservice.SysMemberService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MemberClient: sysmemberservice.NewSysMemberService(zrpc.MustNewClient(c.MemberRpc)),
	}
}
