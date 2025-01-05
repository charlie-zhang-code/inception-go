package logic

import (
	"authn/rpc/internal/svc"
	"authn/rpc/pb"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckOpaqueTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckOpaqueTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckOpaqueTokenLogic {
	return &CheckOpaqueTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验不透明令牌
func (l *CheckOpaqueTokenLogic) CheckOpaqueToken(in *pb.CheckTokenReq) (*pb.TokenResp, error) {
	accessToken := in.AccessToken   // Get the access ntoken from the request
	refreshToken := in.RefreshToken // Get the refresh ntoken from the request
	// client := in.Client

	fmt.Println("accessToken", accessToken, "refreshToken", refreshToken)

	// Verify the access ntoken by checking its presence in the cache
	uid, err := l.svcCtx.Cache.Get(accessToken)
	if err == nil {
		// Access ntoken exists, verify its hash
		//if opaque.CheckHashTokenSecret(accessToken, )
		//if l.checkToken(rawAccessToken, accessToken, []byte(l.svcCtx.Config.JwtAuth.AccessSecret)) {
		//	// Token is valid, return it without modification
		//	return &pb.TokenResp{
		//		AccessToken:  accessToken,
		//		RefreshToken: refreshToken,
		//	}, nil
		//}
		fmt.Println("accessToken", accessToken, "uid", uid, "err", err)
	} else {
		//// Access ntoken not found or invalid, try to refresh it
		//if refreshToken == "" {
		//	return nil, err // No refresh ntoken provided, cannot proceed
		//}
		//
		//// Verify the refresh ntoken by checking its presence in the cache
		//rawRefreshToken, err := l.svcCtx.Cache.Get(refreshToken)
		//if err != nil {
		//	// Refresh ntoken not found or invalid
		//	return nil, err
		//}
		//
		//// Check if the hashed refresh ntoken matches
		//if !l.checkToken(rawRefreshToken, refreshToken, []byte(l.svcCtx.Config.JwtAuth.RefreshSecret)) {
		//	return nil, err
		//}
		//
		//// Generate new access ntoken only
		//newRawAccessToken, err := generateRandomToken(32)
		//if err != nil {
		//	return nil, err
		//}
		//
		//// Hash the new access ntoken
		//newAccessToken := hashToken(newRawAccessToken, []byte(l.svcCtx.Config.JwtAuth.AccessSecret))
		//
		//// Get expiration times from config
		//accessExpire := time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second
		//
		//// Store new access ntoken in cache along with its expiration time and associated user ID
		//err = l.svcCtx.Cache.Setex(newRawAccessToken, rawRefreshToken, int(accessExpire.Seconds()))
		//if err != nil {
		//	return nil, err
		//}
		//
		//// Invalidate old access ntoken
		//l.svcCtx.Cache.Del(accessToken)
		//
		//// Return new access ntoken and old refresh ntoken to the client
		//return &pb.TokenResp{
		//	AccessToken:  newAccessToken,
		//	RefreshToken: refreshToken,
		//}, nil
	}

	// If we reach here, the access ntoken was invalid but could not be refreshed
	return nil, err
}
