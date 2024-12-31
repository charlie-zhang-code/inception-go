package servicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageSysUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPageSysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageSysUserListLogic {
	return &QueryPageSysUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分页列表
func (l *QueryPageSysUserListLogic) QueryPageSysUserList(in *pb.QueryPageSysUserListReq) (*pb.QueryPageSysUserListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryPageSysUserListResp{}, nil
}
