package logic

import (
	"context"
	"convert/field"
	"github.com/jinzhu/copier"
	"status/proto/model"

	"status/rpc/internal/svc"
	"status/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMemberLoginStatusDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberLoginStatusDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLoginStatusDetailLogic {
	return &QueryMemberLoginStatusDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户登录状态详情
func (l *QueryMemberLoginStatusDetailLogic) QueryMemberLoginStatusDetail(in *pb.QueryMemberLoginStatusDetailReq) (*pb.QueryMemberLoginStatusDetailResp, error) {
	var entity model.SysMemberLoginStatus
	result := &pb.QueryMemberLoginStatusDetailResp{}
	err := l.svcCtx.DB.Where("id = ?", in.Id).
		First(&entity).Error

	err = copier.Copy(result, &entity)

	if err := field.TimeFields(&entity, result); err != nil {
		return nil, err
	}

	return &pb.QueryMemberLoginStatusDetailResp{}, err
}
