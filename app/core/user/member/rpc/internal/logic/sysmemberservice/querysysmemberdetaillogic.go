package sysmemberservicelogic

import (
	"context"

	"member/rpc/internal/svc"
	"member/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySysMemberDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySysMemberDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySysMemberDetailLogic {
	return &QuerySysMemberDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户详情
func (l *QuerySysMemberDetailLogic) QuerySysMemberDetail(in *proto.QuerySysMemberDetailReq) (*proto.QuerySysMemberDetailResp, error) {
	// todo: add your logic here and delete this line
	member := l.svcCtx.Member

	res, _ := member.FindOne(l.ctx, in.Id)

	if res == nil {
		return nil, nil
	}

	return &proto.QuerySysMemberDetailResp{
		Id:          res.Id,
		Identify:    res.Identify,
		Username:    res.Username,
		Nickname:    res.Nickname,
		Avatar:      res.Avatar,
		Quote:       res.Quote,
		Birthday:    res.Birthday.String(),
		Gender:      res.Gender,
		Email:       res.Email,
		Telephone:   res.Telephone,
		LoginAt:     res.LoginAt.String(),
		LoginIp:     res.LoginIp,
		LoginRegion: res.LoginRegion,
		LoginOs:     res.LoginOs,
		Status:      res.Status,
		Remark:      res.Remark,
		CreateAt:    res.CreateAt.String(),
		CreateBy:    res.CreateBy,
		UpdateAt:    res.UpdateAt.String(),
		UpdateBy:    res.UpdateBy,
	}, nil
}
