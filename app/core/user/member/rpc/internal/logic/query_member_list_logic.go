package logic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberListLogic {
	return &QueryMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户所有列表
func (l *QueryMemberListLogic) QueryMemberList(in *pb.QueryMemberListReq) (*pb.QueryMemberListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryMemberListResp{}, nil
}
