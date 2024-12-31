package servicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserLogic {
	return &UpdateSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户
func (l *UpdateSysUserLogic) UpdateSysUser(in *pb.UpdateSysUserReq) (*pb.UpdateSysUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateSysUserResp{}, nil
}
