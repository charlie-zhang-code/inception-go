package logic

import (
	"context"

	"group/rpc/internal/svc"
	"group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryKeywordPageGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryKeywordPageGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryKeywordPageGroupListLogic {
	return &QueryKeywordPageGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户组所有列表
func (l *QueryKeywordPageGroupListLogic) QueryKeywordPageGroupList(in *pb.QueryKeywordPageGroupListReq) (*pb.QueryKeywordPageGroupListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryKeywordPageGroupListResp{}, nil
}
