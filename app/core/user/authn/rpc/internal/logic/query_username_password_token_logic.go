package logic

import (
	"authn/rpc/internal/svc"
	"authn/rpc/pb"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type QueryUsernamePasswordTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUsernamePasswordTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUsernamePasswordTokenLogic {
	return &QueryUsernamePasswordTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过用户名密码获取token
func (l *QueryUsernamePasswordTokenLogic) QueryUsernamePasswordToken(in *pb.UsernamePasswordTokenReq) (*pb.TokenResp, error) {

	// 临时测试
	uid := "sdgwsectwerg-sagrcwrg-cadevaw"

	now := time.Now().Unix()

	// Access Token
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := GetAccessToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, uid)
	if err != nil {
		return nil, err
	}

	// Refresh Token
	refreshExpire := l.svcCtx.Config.JwtAuth.RefreshExpire
	refreshToken, err := GetRefreshToken(l.svcCtx.Config.JwtAuth.RefreshSecret, now, refreshExpire, uid)
	if err != nil {
		return nil, err
	}

	expirationTime := now + int64(accessExpire)

	return &pb.TokenResp{
		AccessToken:  accessToken,
		AccessExpire: strconv.FormatInt(expirationTime, 10),
		RefreshAfter: strconv.FormatInt(expirationTime-time.Now().Unix(), 10),
		RefreshToken: refreshToken,
	}, nil
}

func GetAccessToken(secretKey string, iat int64, expireSeconds int, uid string) (string, error) {
	claims := jwt.MapClaims{
		"exp": iat + int64(expireSeconds),
		"iat": iat,
		"uid": uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func GetRefreshToken(secretKey string, iat int64, expireSeconds int, uid string) (string, error) {
	claims := jwt.MapClaims{
		"exp": iat + int64(expireSeconds),
		"iat": iat,
		"uid": uid,
		"typ": "refresh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
