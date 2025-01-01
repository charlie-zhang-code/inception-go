package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSysMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysMemberLogic {
	return &DeleteSysMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteSysMemberLogic) DeleteSysMember(in *proto.DeleteSysMemberReq) (*proto.DeleteSysMemberResp, error) {
	// todo: add your logic here and delete this line

	return &proto.DeleteSysMemberResp{}, nil
}
