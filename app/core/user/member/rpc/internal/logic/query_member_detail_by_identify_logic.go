package logic

import (
	"context"
	"convert/field"
	"github.com/jinzhu/copier"
	"member/proto/model"

	"member/rpc/internal/svc"
	"member/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberDetailByIdentifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberDetailByIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberDetailByIdentifyLogic {
	return &QueryMemberDetailByIdentifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情（通过id）
func (l *QueryMemberDetailByIdentifyLogic) QueryMemberDetailByIdentify(in *pb.QueryMemberDetailByIdentifyReq) (*pb.QueryMemberDetailByIdentifyResp, error) {
	var entity model.SysMember
	result := &pb.QueryMemberDetailByIdentifyResp{}
	err := l.svcCtx.DB.Where("identify = ?", in.Identify).
		First(&entity).Error

	err = copier.Copy(result, &entity)

	if err := field.TimeFields(&entity, result); err != nil {
		return nil, err
	}

	return result, err
}
