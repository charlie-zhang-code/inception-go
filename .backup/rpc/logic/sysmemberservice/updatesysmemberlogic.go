package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysMemberLogic {
	return &UpdateSysMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户
func (l *UpdateSysMemberLogic) UpdateSysMember(in *proto.UpdateSysMemberReq) (*proto.UpdateSysMemberResp, error) {
	// todo: add your logic here and delete this line

	return &proto.UpdateSysMemberResp{}, nil
}
