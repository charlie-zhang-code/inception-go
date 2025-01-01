package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysMemberLogic {
	return &AddSysMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysMemberLogic) AddSysMember(req *types.AddSysMemberReq) (resp *types.AddSysMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
