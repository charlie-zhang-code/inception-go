package logic

import (
	"context"

	"group/rpc/internal/svc"
	"group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPageGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageGroupListLogic {
	return &QueryPageGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户组分页列表
func (l *QueryPageGroupListLogic) QueryPageGroupList(in *pb.QueryPageGroupListReq) (*pb.QueryPageGroupListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryPageGroupListResp{}, nil
}
