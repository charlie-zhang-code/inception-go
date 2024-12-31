package servicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserStatusLogic {
	return &UpdateSysUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户状态
func (l *UpdateSysUserStatusLogic) UpdateSysUserStatus(in *pb.UpdateSysUserStatusReq) (*pb.UpdateSysUserStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateSysUserStatusResp{}, nil
}
