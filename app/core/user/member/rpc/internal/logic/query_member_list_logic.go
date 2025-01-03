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

type QueryMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberListLogic {
	return &QueryMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户所有列表
func (l *QueryMemberListLogic) QueryMemberList(in *pb.QueryMemberListReq) (*pb.QueryMemberListResp, error) {
	db := l.svcCtx.DB.Model(&model.SysMember{})

	var resultList []model.SysMember

	err := db.
		Where("deleted = ?", 0).
		Find(&resultList).Error

	result := &pb.QueryMemberListResp{
		List: make([]*pb.MemberListData, 0, len(resultList)),
	}

	for _, item := range resultList {
		res := &pb.MemberListData{}
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
