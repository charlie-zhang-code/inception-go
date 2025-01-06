package logic

import (
	"context"

	"group/rpc/internal/svc"
	"group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户组
func (l *AddGroupLogic) AddGroup(in *pb.AddGroupReq) (*pb.AddGroupResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddGroupResp{}, nil
}
