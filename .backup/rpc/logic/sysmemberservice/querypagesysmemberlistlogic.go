package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageSysMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPageSysMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageSysMemberListLogic {
	return &QueryPageSysMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分页列表
func (l *QueryPageSysMemberListLogic) QueryPageSysMemberList(in *proto.QueryPageSysMemberListReq) (*proto.QueryPageSysMemberListResp, error) {
	// todo: add your logic here and delete this line

	return &proto.QueryPageSysMemberListResp{}, nil
}
