package logic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberDetailByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberDetailByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberDetailByIdLogic {
	return &GetMemberDetailByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情（通过id）
func (l *GetMemberDetailByIdLogic) GetMemberDetailById(in *pb.QueryMemberDetailById) (*pb.MemberData, error) {
	// todo: add your logic here and delete this line

	return &pb.MemberData{}, nil
}
