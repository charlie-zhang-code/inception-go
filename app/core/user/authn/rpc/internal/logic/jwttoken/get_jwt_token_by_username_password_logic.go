package jwttokenlogic

import (
	"authn/rpc/internal/svc"
	"authn/rpc/pb"
	"context"
	"golang.org/x/crypto/bcrypt"
	"member/rpc/member"
	"ntoken"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJwtTokenByUsernamePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJwtTokenByUsernamePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJwtTokenByUsernamePasswordLogic {
	return &GetJwtTokenByUsernamePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过用户名密码获取 jwt token
func (l *GetJwtTokenByUsernamePasswordLogic) GetJwtTokenByUsernamePassword(in *pb.UsernamePasswordReq) (*pb.JwtTokenResp, error) {
	var accessSecret = l.svcCtx.Config.JwtAuth.AccessSecret
	var accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire
	var RefreshSecret = l.svcCtx.Config.JwtAuth.RefreshSecret
	var refreshExpire = l.svcCtx.Config.JwtAuth.RefreshExpire

	detail, err := l.svcCtx.MemberService.QueryMemberDetailByUsername(l.ctx, &member.QueryMemberDetailByUsernameReq{
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(detail.Password), []byte(in.Password))
	if err != nil {
		return &pb.JwtTokenResp{}, err // 密码不匹配
	}

	now := time.Now().Unix()

	accessToken, _ := ntoken.GenerateJwt(now, detail.Identify, accessSecret, nil, accessExpire)
	refreshToken, _ := ntoken.GenerateJwt(now, detail.Identify, RefreshSecret, nil, refreshExpire)

	return &pb.JwtTokenResp{
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
		RefreshAfter:  time.Now().Add(time.Duration(refreshExpire) * time.Second).Format(time.DateTime),
		AccessExpire:  time.Now().Add(time.Duration(accessExpire) * time.Second).Format(time.DateTime),
		RefreshExpire: time.Now().Add(time.Duration(refreshExpire) * time.Second).Format(time.DateTime),
		Client:        in.Client,
	}, nil
}
