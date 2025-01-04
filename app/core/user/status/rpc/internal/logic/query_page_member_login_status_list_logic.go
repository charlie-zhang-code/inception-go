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

type QueryPageMemberLoginStatusListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPageMemberLoginStatusListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageMemberLoginStatusListLogic {
	return &QueryPageMemberLoginStatusListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户登录状态分页列表
func (l *QueryPageMemberLoginStatusListLogic) QueryPageMemberLoginStatusList(in *pb.QueryPageMemberLoginStatusListReq) (*pb.QueryPageMemberLoginStatusListResp, error) {
	db := l.svcCtx.DB.Model(&model.SysMemberLoginStatus{})

	var resultList []model.SysMemberLoginStatus
	var total int64
	err := db.Count(&total).Error
	err = db.Where("deleted = ?", 0).
		Limit(int(in.PageSize)).
		Offset(int(in.PageSize * (in.PageNum - 1))).
		Find(&resultList).Error

	result := &pb.QueryPageMemberLoginStatusListResp{
		List:  make([]*pb.PageMemberLoginStatusListData, 0, len(resultList)),
		Total: total,
	}

	for _, item := range resultList {
		res := &pb.PageMemberLoginStatusListData{}
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
