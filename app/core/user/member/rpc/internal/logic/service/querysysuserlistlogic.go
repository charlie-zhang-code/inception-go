package servicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysUserListLogic {
	return &QuerySysUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户所有列表
func (l *QuerySysUserListLogic) QuerySysUserList(in *pb.QuerySysUserListReq) (*pb.QuerySysUserListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QuerySysUserListResp{}, nil
}
