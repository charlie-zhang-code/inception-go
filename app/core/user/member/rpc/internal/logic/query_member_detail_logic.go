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

type QueryMemberDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberDetailLogic {
	return &QueryMemberDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情（通过id）
func (l *QueryMemberDetailLogic) QueryMemberDetail(in *pb.QueryMemberDetailReq) (*pb.QueryMemberDetailResp, error) {
	var entity model.SysMember
	result := &pb.QueryMemberDetailResp{}
	err := l.svcCtx.DB.Where("id = ?", in.Id).
		First(&entity).Error

	err = copier.Copy(result, &entity)

	if err := field.TimeFields(&entity, result); err != nil {
		return nil, err
	}

	return result, err
}
