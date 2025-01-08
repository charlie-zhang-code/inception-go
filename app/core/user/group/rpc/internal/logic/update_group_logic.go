package logic

import (
	"context"

	"group/rpc/internal/svc"
	"group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户组
func (l *UpdateGroupLogic) UpdateGroup(in *pb.UpdateGroupReq) (*pb.UpdateGroupResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateGroupResp{}, nil
}
