package logic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberDetailByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberDetailByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberDetailByUsernameLogic {
	return &GetMemberDetailByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情（通过用户名）
func (l *GetMemberDetailByUsernameLogic) GetMemberDetailByUsername(in *pb.QueryMemberDetailByUsername) (*pb.MemberData, error) {
	// todo: add your logic here and delete this line

	return &pb.MemberData{}, nil
}
