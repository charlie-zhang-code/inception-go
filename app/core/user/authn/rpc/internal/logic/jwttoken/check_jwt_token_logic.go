package jwttokenlogic

import (
	"context"
	"member/rpc/member"
	"ntoken"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckJwtTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckJwtTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckJwtTokenLogic {
	return &CheckJwtTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验 jwt token
func (l *CheckJwtTokenLogic) CheckJwtToken(in *pb.CheckJwtTokenReq) (*pb.CheckJwtTokenResp, error) {
	claims, err := ntoken.ParseJwt(in.Token, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return nil, err
	}

	detail, err := l.svcCtx.MemberService.QueryMemberDetailByIdentify(l.ctx, &member.QueryMemberDetailByIdentifyReq{
		Identify: claims.Uid,
	})
	if err != nil {
		return &pb.CheckJwtTokenResp{
			Success: false,
		}, err
	}

	if detail.Identify != claims.Uid {
		return &pb.CheckJwtTokenResp{
			Success: false,
		}, err
	}

	return &pb.CheckJwtTokenResp{
		Success: true,
	}, nil
}
