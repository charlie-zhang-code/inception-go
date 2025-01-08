package logic

import (
	"context"

	"role/rpc/internal/svc"
	"role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除角色
func (l *DeleteRoleLogic) DeleteRole(in *pb.DeleteRoleReq) (*pb.DeleteRoleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteRoleResp{}, nil
}
