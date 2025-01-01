package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysMemberStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysMemberStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysMemberStatusLogic {
	return &UpdateSysMemberStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户状态
func (l *UpdateSysMemberStatusLogic) UpdateSysMemberStatus(in *proto.UpdateSysMemberStatusReq) (*proto.UpdateSysMemberStatusResp, error) {
	// todo: add your logic here and delete this line

	return &proto.UpdateSysMemberStatusResp{}, nil
}
