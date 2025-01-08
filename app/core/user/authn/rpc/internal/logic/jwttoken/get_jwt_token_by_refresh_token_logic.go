package jwttokenlogic

import (
	"context"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJwtTokenByRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJwtTokenByRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJwtTokenByRefreshTokenLogic {
	return &GetJwtTokenByRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过刷新令牌获取 jwt token
func (l *GetJwtTokenByRefreshTokenLogic) GetJwtTokenByRefreshToken(in *pb.JwtRefreshTokenReq) (*pb.JwtTokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.JwtTokenResp{}, nil
}
