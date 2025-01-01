package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageSysMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageSysMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageSysMemberListLogic {
	return &QueryPageSysMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageSysMemberListLogic) QueryPageSysMemberList(req *types.QueryPageSysMemberListReq) (resp *types.QueryPageSysMemberListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
