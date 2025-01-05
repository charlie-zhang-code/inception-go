package logic

import (
	"authn/rpc/internal/svc"
	"authn/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"ntoken/opaque"
	"strconv"
	"time"
)

type QueryUsernamePasswordOpaqueTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUsernamePasswordOpaqueTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUsernamePasswordOpaqueTokenLogic {
	return &QueryUsernamePasswordOpaqueTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 不透明令牌颁发
func (l *QueryUsernamePasswordOpaqueTokenLogic) QueryUsernamePasswordOpaqueToken(in *pb.UsernamePasswordTokenReq) (*pb.TokenResp, error) {
	userID, err := l.authenticateUser(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	// Generate raw opaque tokens
	rawAccessToken, err := opaque.GenerateRandomToken(32) // 32 bytes for the raw access ntoken
	if err != nil {
		return nil, err
	}
	rawRefreshToken, err := opaque.GenerateRandomToken(32)
	if err != nil {
		return nil, err
	}

	// Hash the tokens with their respective secrets
	accessToken, _ := opaque.HashToken(rawAccessToken, []byte(l.svcCtx.Config.JwtAuth.AccessSecret))
	refreshToken, _ := opaque.HashToken(rawRefreshToken, []byte(l.svcCtx.Config.JwtAuth.RefreshSecret))

	// Get expiration times from config
	accessExpire := time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second
	refreshExpire := time.Duration(l.svcCtx.Config.JwtAuth.RefreshExpire) * time.Second

	// Calculate expiration times
	expireTime := time.Now()
	accessExpireTime := expireTime.Add(accessExpire)
	refreshExpireTime := expireTime.Add(refreshExpire)

	// Store tokens in cache along with their expiration times and associated user ID
	err = l.svcCtx.Cache.Setex(rawAccessToken, userID, int(accessExpire.Seconds()))
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Cache.Setex(rawRefreshToken, userID, int(refreshExpire.Seconds()))
	if err != nil {
		return nil, err
	}

	// Return hashed tokens to the client
	return &pb.TokenResp{
		AccessToken:  accessToken,
		AccessExpire: strconv.FormatInt(accessExpireTime.Unix(), 10),
		RefreshToken: refreshToken, // Refresh in client
		RefreshAfter: strconv.FormatInt(refreshExpireTime.Unix()-time.Now().Unix(), 10),
	}, nil
}

func (l *QueryUsernamePasswordOpaqueTokenLogic) authenticateUser(username, password string) (string, error) {
	// todo: Implement your user authentication logic here.
	// For now, let's assume it always returns "user123".
	return "user123", nil
}
