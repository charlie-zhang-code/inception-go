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

type QueryKeywordMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryKeywordMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryKeywordMemberListLogic {
	return &QueryKeywordMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询关键字用户列表
func (l *QueryKeywordMemberListLogic) QueryKeywordMemberList(in *pb.QueryKeywordMemberListReq) (*pb.QueryKeywordMemberListResp, error) {
	db := l.svcCtx.DB.Model(&model.SysMember{})

	var resultList []model.SysMember

	keyword := "%" + in.Keyword + "%"
	err := db.
		Where("deleted = ?", 0).
		Where("nickname LIKE ? OR email LIKE ? OR telephone LIKE ?", keyword, keyword, keyword).
		Find(&resultList).Error

	result := &pb.QueryKeywordMemberListResp{
		List: make([]*pb.KeywordMemberListData, 0, len(resultList)),
	}

	for _, item := range resultList {
		res := &pb.KeywordMemberListData{}
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
