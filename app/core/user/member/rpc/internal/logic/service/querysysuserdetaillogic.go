package servicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysUserDetailLogic {
	return &QuerySysUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情
func (l *QuerySysUserDetailLogic) QuerySysUserDetail(in *pb.QuerySysUserDetailReq) (*pb.QuerySysUserDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QuerySysUserDetailResp{}, nil
}
