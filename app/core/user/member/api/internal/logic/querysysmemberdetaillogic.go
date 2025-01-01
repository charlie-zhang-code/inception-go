package logic

import (
	"context"
	"member/rpc/proto"

	"member/api/internal/svc"
	"member/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysMemberDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySysMemberDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysMemberDetailLogic {
	return &QuerySysMemberDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySysMemberDetailLogic) QuerySysMemberDetail(req *types.QuerySysMemberDetailReq) (resp *types.QuerySysMemberDetailResp, err error) {
	// todo: add your logic here and delete this line
	res, _ := l.svcCtx.MemberClient.QuerySysMemberDetail(l.ctx, &proto.QuerySysMemberDetailReq{
		Id: req.Id,
	})

	if res == nil {
		return nil, nil
	}

	data := types.QuerySysMemberDetailData{
		Id:          res.Id,
		Identify:    res.Identify,
		Username:    res.Username,
		Nickname:    res.Nickname,
		Avatar:      res.Avatar,
		Quote:       res.Quote,
		Birthday:    res.Birthday,
		Gender:      res.Gender,
		Email:       res.Email,
		Telephone:   res.Telephone,
		LoginAt:     res.LoginAt,
		LoginIp:     res.LoginIp,
		LoginRegion: res.LoginRegion,
		LoginOs:     res.LoginOs,
		Status:      res.Status,
		Remark:      res.Remark,
		CreateAt:    res.CreateAt,
		CreateBy:    res.CreateBy,
		UpdateAt:    res.UpdateAt,
		UpdateBy:    res.UpdateBy,
	}

	return &types.QuerySysMemberDetailResp{
		Code:    "200",
		Message: "OK",
		Data:    data,
	}, nil
}
