package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysMemberListLogic {
	return &QuerySysMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户所有列表
func (l *QuerySysMemberListLogic) QuerySysMemberList(in *proto.QuerySysMemberListReq) (*proto.QuerySysMemberListResp, error) {
	// todo: add your logic here and delete this line

	return &proto.QuerySysMemberListResp{}, nil
}
