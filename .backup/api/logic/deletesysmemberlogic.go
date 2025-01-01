package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysMemberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysMemberLogic {
	return &DeleteSysMemberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysMemberLogic) DeleteSysMember(req *types.DeleteSysMemberReq) (resp *types.DeleteSysMemberResp, err error) {
	// todo: add your logic here and delete this line

	return
}
