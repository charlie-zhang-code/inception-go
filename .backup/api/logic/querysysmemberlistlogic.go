package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysMemberListLogic {
	return &QuerySysMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysMemberListLogic) QuerySysMemberList(req *types.QuerySysMemberListReq) (resp *types.QuerySysMemberListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
