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

type QueryPageMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPageMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPageMemberListLogic {
	return &QueryPageMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分页列表
func (l *QueryPageMemberListLogic) QueryPageMemberList(in *pb.QueryPageMemberListReq) (*pb.QueryPageMemberListResp, error) {
	db := l.svcCtx.DB.Model(&model.SysMember{})

	var resultList []model.SysMember
	var total int64
	err := db.Count(&total).Error
	err = db.Where("deleted = ?", 0).
		Limit(int(in.PageSize)).
		Offset(int(in.PageSize * (in.PageNum - 1))).
		Find(&resultList).Error

	result := &pb.QueryPageMemberListResp{
		List:  make([]*pb.PageMemberListData, 0, len(resultList)),
		Total: total,
	}

	for _, item := range resultList {
		res := &pb.PageMemberListData{}
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
