package authenticationservicelogic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"authn/rpc/internal/svc"
	"authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsernamePasswordTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsernamePasswordTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsernamePasswordTokenLogic {
	return &UsernamePasswordTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过用户名密码获取token
func (l *UsernamePasswordTokenLogic) UsernamePasswordToken(in *pb.UsernamePasswordTokenReq) (*pb.UsernamePasswordTokenResp, error) {
	// todo: add your logic here and delete this line

	fmt.Println("Username: " + in.Username)
	fmt.Println("Password: " + in.Password)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))

	resd := bcrypt.CompareHashAndPassword([]byte(string(hashedPassword)), []byte(in.Password))
	fmt.Println(resd)

	var accessExpire = l.svcCtx.Config.JwtAuth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	return &pb.UsernamePasswordTokenResp{
		AccessToken:  accessToken,
		AccessExpire: strconv.FormatInt(now+accessExpire, 10),
		RefreshAfter: strconv.FormatInt(now+accessExpire/2, 10),
	}, nil
}

func (l *UsernamePasswordTokenLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
