package logic

import (
	"context"

	"role/rpc/internal/svc"
	"role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色状态
func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *pb.UpdateRoleStatusReq) (*pb.UpdateRoleStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateRoleStatusResp{}, nil
}
