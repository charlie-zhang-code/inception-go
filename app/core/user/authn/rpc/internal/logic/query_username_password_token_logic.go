package logic

import (
	"authn/rpc/internal/svc"
	"authn/rpc/pb"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, 145)
	if err != nil {
		return nil, err
	}

	return &pb.TokenResp{
		AccessToken:  accessToken,
		AccessExpire: strconv.FormatInt(accessExpire, 10),
	}, nil
}

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
