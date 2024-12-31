package logic

import (
	"context"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysUserDetailLogic {
	return &QuerySysUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysUserDetailLogic) QuerySysUserDetail(req *types.QuerySysUserDetailReq) (resp *types.QuerySysUserDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
