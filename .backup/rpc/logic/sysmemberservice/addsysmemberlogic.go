package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysMemberLogic {
	return &AddSysMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户
func (l *AddSysMemberLogic) AddSysMember(in *proto.AddSysMemberReq) (*proto.AddSysMemberResp, error) {
	// todo: add your logic here and delete this line

	return &proto.AddSysMemberResp{}, nil
}
