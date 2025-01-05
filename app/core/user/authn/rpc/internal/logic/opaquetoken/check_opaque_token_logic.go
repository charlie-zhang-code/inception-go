package opaquetokenlogic

import (
	"context"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOpaqueTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckOpaqueTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOpaqueTokenLogic {
	return &CheckOpaqueTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验 opaque token
func (l *CheckOpaqueTokenLogic) CheckOpaqueToken(in *pb.CheckOpaqueTokenReq) (*pb.OpaqueTokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OpaqueTokenResp{}, nil
}
