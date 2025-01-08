package opaquetokenlogic

import (
	"context"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOpaqueTokenByRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOpaqueTokenByRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOpaqueTokenByRefreshTokenLogic {
	return &GetOpaqueTokenByRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过刷新令牌获取 opaque token
func (l *GetOpaqueTokenByRefreshTokenLogic) GetOpaqueTokenByRefreshToken(in *pb.OpaqueRefreshTokenReq) (*pb.OpaqueTokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OpaqueTokenResp{}, nil
}
