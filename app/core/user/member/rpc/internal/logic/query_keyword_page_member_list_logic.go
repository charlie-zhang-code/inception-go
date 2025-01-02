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

type QueryKeywordPageMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryKeywordPageMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryKeywordPageMemberListLogic {
	return &QueryKeywordPageMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询关键字用户分页列表
func (l *QueryKeywordPageMemberListLogic) QueryKeywordPageMemberList(in *pb.QueryKeywordPageMemberListReq) (*pb.QueryKeywordPageMemberListResp, error) {
	db := l.svcCtx.DB.Model(&model.SysMember{})

	var resultList []model.SysMember
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, err
	}

	keyword := "%" + in.Keyword + "%"
	err = db.Where("nickname LIKE ? OR email LIKE ?", keyword, keyword).Where("deleted = ?", 0).Limit(int(in.PageSize)).Offset(int(in.PageSize * (in.PageNum - 1))).Find(&resultList).Error

	result := &pb.QueryKeywordPageMemberListResp{
		List:  make([]*pb.KeywordPageMemberListData, 0, len(resultList)),
		Total: total,
	}

	for _, item := range resultList {
		res := &pb.KeywordPageMemberListData{}
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
