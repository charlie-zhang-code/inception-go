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

type QueryMemberLoginStatusListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberLoginStatusListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLoginStatusListLogic {
	return &QueryMemberLoginStatusListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户登录状态所有列表
func (l *QueryMemberLoginStatusListLogic) QueryMemberLoginStatusList(in *pb.QueryMemberLoginStatusListReq) (*pb.QueryMemberLoginStatusListResp, error) {
	db := l.svcCtx.DB.Model(&model.SysMemberLoginStatus{})

	var resultList []model.SysMemberLoginStatus

	err := db.
		Where("deleted = ?", 0).
		Find(&resultList).Error

	result := &pb.QueryMemberLoginStatusListResp{
		List: make([]*pb.MemberLoginStatusListData, 0, len(resultList)),
	}

	for _, item := range resultList {
		res := &pb.MemberLoginStatusListData{}
		if err := copier.Copy(res, &item); err != nil {
			return nil, err
		}

		if err := field.TimeFields(&item, res); err != nil {
			return nil, err
		}

		result.List = append(result.List, res)
	}

	return result, err
}
