package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysUserLogic {
	return &DeleteSysUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysUserLogic) DeleteSysUser(req *types.DeleteSysUserReq) (resp *types.DeleteSysUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
