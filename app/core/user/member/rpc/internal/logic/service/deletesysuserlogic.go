package servicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysUserLogic {
	return &DeleteSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteSysUserLogic) DeleteSysUser(in *pb.DeleteSysUserReq) (*pb.DeleteSysUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteSysUserResp{}, nil
}
