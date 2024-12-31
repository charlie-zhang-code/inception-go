package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysUserStatusLogic {
	return &UpdateSysUserStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysUserStatusLogic) UpdateSysUserStatus(req *types.UpdateSysUserStatusReq) (resp *types.UpdateSysUserStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
