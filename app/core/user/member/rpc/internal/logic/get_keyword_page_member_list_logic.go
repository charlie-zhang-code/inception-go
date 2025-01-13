package logic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKeywordPageMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetKeywordPageMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKeywordPageMemberListLogic {
	return &GetKeywordPageMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分页列表
func (l *GetKeywordPageMemberListLogic) GetKeywordPageMemberList(in *pb.QueryPageMemberList) (*pb.QueryPageMemberListResult, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryPageMemberListResult{}, nil
}
