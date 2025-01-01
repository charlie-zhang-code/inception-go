package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysMemberLogic {
	return &UpdateSysMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysMemberLogic) UpdateSysMember(req *types.UpdateSysMemberReq) (resp *types.UpdateSysMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
