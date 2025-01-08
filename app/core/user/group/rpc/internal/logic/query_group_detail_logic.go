package logic

import (
	"context"

	"group/rpc/internal/svc"
	"group/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryGroupDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryGroupDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryGroupDetailLogic {
	return &QueryGroupDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户组详情
func (l *QueryGroupDetailLogic) QueryGroupDetail(in *pb.QueryGroupDetailReq) (*pb.QueryGroupDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.QueryGroupDetailResp{}, nil
}
