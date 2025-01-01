package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysMemberStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysMemberStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysMemberStatusLogic {
	return &UpdateSysMemberStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysMemberStatusLogic) UpdateSysMemberStatus(req *types.UpdateSysMemberStatusReq) (resp *types.UpdateSysMemberStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
