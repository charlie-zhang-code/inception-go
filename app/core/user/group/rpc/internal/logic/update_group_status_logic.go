package logic

import (
	"context"

	"group/rpc/internal/svc"
	"group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupStatusLogic {
	return &UpdateGroupStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户组状态
func (l *UpdateGroupStatusLogic) UpdateGroupStatus(in *pb.UpdateGroupStatusReq) (*pb.UpdateGroupStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateGroupStatusResp{}, nil
}
