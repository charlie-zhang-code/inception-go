package logic

import (
	"context"

	"role/rpc/internal/svc"
	"role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色详情
func (l *QueryRoleDetailLogic) QueryRoleDetail(in *pb.QueryRoleDetailReq) (*pb.QueryRoleDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryRoleDetailResp{}, nil
}
