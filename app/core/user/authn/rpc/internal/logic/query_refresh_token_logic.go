package logic

import (
	"context"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRefreshTokenLogic {
	return &QueryRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 刷新令牌
func (l *QueryRefreshTokenLogic) QueryRefreshToken(in *pb.RefreshTokenReq) (*pb.TokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.TokenResp{}, nil
}
