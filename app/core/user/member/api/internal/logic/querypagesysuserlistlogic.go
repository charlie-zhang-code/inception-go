package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPageSysUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPageSysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageSysUserListLogic {
	return &QueryPageSysUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPageSysUserListLogic) QueryPageSysUserList(req *types.QueryPageSysUserListReq) (resp *types.QueryPageSysUserListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
