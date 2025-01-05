package opaquetokenlogic

import (
	"context"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOpaqueTokenByUsernamePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOpaqueTokenByUsernamePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOpaqueTokenByUsernamePasswordLogic {
	return &GetOpaqueTokenByUsernamePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过用户名密码获取 opaque token
func (l *GetOpaqueTokenByUsernamePasswordLogic) GetOpaqueTokenByUsernamePassword(in *pb.UsernamePasswordReq) (*pb.OpaqueTokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OpaqueTokenResp{}, nil
}
