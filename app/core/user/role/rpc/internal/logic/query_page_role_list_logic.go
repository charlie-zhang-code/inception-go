package logic

import (
	"context"

	"role/rpc/internal/svc"
	"role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPageRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageRoleListLogic {
	return &QueryPageRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色分页列表
func (l *QueryPageRoleListLogic) QueryPageRoleList(in *pb.QueryPageRoleListReq) (*pb.QueryPageRoleListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryPageRoleListResp{}, nil
}
