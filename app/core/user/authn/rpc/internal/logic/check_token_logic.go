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

type CheckTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验Token
func (l *CheckTokenLogic) CheckToken(in *pb.CheckTokenReq) (*pb.TokenResp, error) {
	var _ jwt.MapClaims
	token, err := jwt.Parse(in.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	if err != nil {
		// 如果解析失败，尝试用刷新令牌获取新的访问令牌
		return l.refreshToken(in)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := claims["exp"].(float64)
		if time.Now().Unix() < int64(expirationTime) {
			// Token 未过期，返回原始信息
			return &pb.TokenResp{
				AccessToken:  in.AccessToken,
				AccessExpire: strconv.FormatInt(int64(expirationTime), 10),
				RefreshAfter: strconv.FormatInt(int64(expirationTime)-time.Now().Unix(), 10),
				RefreshToken: in.RefreshToken,
			}, nil
		} else {
			// Token 已过期，尝试用刷新令牌获取新的访问令牌
			return l.refreshToken(in)
		}
	}

	return nil, err
}

func (l *CheckTokenLogic) refreshToken(in *pb.CheckTokenReq) (*pb.TokenResp, error) {
	token, err := jwt.Parse(in.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(l.svcCtx.Config.JwtAuth.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if typ, exists := claims["typ"]; exists && typ == "refresh" {
			uid := claims["uid"].(string)

			now := time.Now().Unix()

			// Access Token
			accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
			accessToken, err := GetAccessToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, uid)
			if err != nil {
				return nil, err
			}

			// Refresh Token (可选：根据需求决定是否生成新的刷新令牌)
			refreshToken := in.RefreshToken // 使用旧的刷新令牌
			/*
				refreshExpire := l.svcCtx.Config.JwtAuth.RefreshExpire
				refreshToken, err = GetRefreshToken(l.svcCtx.Config.JwtAuth.RefreshSecret, now, refreshExpire, uid)
				if err != nil {
					return nil, err
				}
			*/

			expirationTime := now + int64(accessExpire)

			return &pb.TokenResp{
				AccessToken:  accessToken,
				AccessExpire: strconv.FormatInt(expirationTime, 10),
				RefreshAfter: strconv.FormatInt(expirationTime-now, 10),
				RefreshToken: refreshToken,
			}, nil
		}
	}

	return nil, err
}
