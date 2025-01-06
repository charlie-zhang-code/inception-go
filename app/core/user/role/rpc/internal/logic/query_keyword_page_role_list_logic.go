package logic

import (
	"context"

	"role/rpc/internal/svc"
	"role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryKeywordPageRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryKeywordPageRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryKeywordPageRoleListLogic {
	return &QueryKeywordPageRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询关键字角色分页列表
func (l *QueryKeywordPageRoleListLogic) QueryKeywordPageRoleList(in *pb.QueryKeywordPageRoleListReq) (*pb.QueryKeywordPageRoleListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryKeywordPageRoleListResp{}, nil
}
