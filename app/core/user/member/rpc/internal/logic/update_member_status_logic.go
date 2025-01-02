package logic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberStatusLogic {
	return &UpdateMemberStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户状态
func (l *UpdateMemberStatusLogic) UpdateMemberStatus(in *pb.UpdateMemberStatusReq) (*pb.UpdateMemberStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMemberStatusResp{}, nil
}
