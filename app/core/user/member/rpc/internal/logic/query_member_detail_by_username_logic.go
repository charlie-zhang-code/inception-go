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

type QueryMemberDetailByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberDetailByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberDetailByUsernameLogic {
	return &QueryMemberDetailByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情（通过用户名）
func (l *QueryMemberDetailByUsernameLogic) QueryMemberDetailByUsername(in *pb.QueryMemberDetailByUsernameReq) (*pb.QueryMemberDetailByUsernameResp, error) {
	var entity model.SysMember
	result := &pb.QueryMemberDetailByUsernameResp{}
	err := l.svcCtx.DB.Where("username = ?", in.Username).
		First(&entity).Error

	err = copier.Copy(result, &entity)

	if err := field.TimeFields(&entity, result); err != nil {
		return nil, err
	}

	return result, err
}
