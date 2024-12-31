package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysUserListLogic {
	return &QuerySysUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysUserListLogic) QuerySysUserList(req *types.QuerySysUserListReq) (resp *types.QuerySysUserListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
