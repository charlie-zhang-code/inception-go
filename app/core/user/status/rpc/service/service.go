// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: main.proto

package service

import (
	"context"

	"status/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddMemberLoginStatusReq            = pb.AddMemberLoginStatusReq
	AddMemberLoginStatusResp           = pb.AddMemberLoginStatusResp
	MemberLoginStatusListData          = pb.MemberLoginStatusListData
	PageMemberLoginStatusListData      = pb.PageMemberLoginStatusListData
	QueryMemberLoginStatusDetailReq    = pb.QueryMemberLoginStatusDetailReq
	QueryMemberLoginStatusDetailResp   = pb.QueryMemberLoginStatusDetailResp
	QueryMemberLoginStatusListReq      = pb.QueryMemberLoginStatusListReq
	QueryMemberLoginStatusListResp     = pb.QueryMemberLoginStatusListResp
	QueryPageMemberLoginStatusListReq  = pb.QueryPageMemberLoginStatusListReq
	QueryPageMemberLoginStatusListResp = pb.QueryPageMemberLoginStatusListResp
	UpdateMemberLogoutStatusReq        = pb.UpdateMemberLogoutStatusReq
	UpdateMemberLogoutStatusResp       = pb.UpdateMemberLogoutStatusResp

	Service interface {
		// 添加用户登录状态
		AddMemberLoginStatus(ctx context.Context, in *AddMemberLoginStatusReq, opts ...grpc.CallOption) (*AddMemberLoginStatusResp, error)
		// 更新用户登录状态
		UpdateMemberLogoutStatus(ctx context.Context, in *UpdateMemberLogoutStatusReq, opts ...grpc.CallOption) (*UpdateMemberLogoutStatusResp, error)
		// 查询用户登录状态详情
		QueryMemberLoginStatusDetail(ctx context.Context, in *QueryMemberLoginStatusDetailReq, opts ...grpc.CallOption) (*QueryMemberLoginStatusDetailResp, error)
		// 查询用户登录状态分页列表
		QueryPageMemberLoginStatusList(ctx context.Context, in *QueryPageMemberLoginStatusListReq, opts ...grpc.CallOption) (*QueryPageMemberLoginStatusListResp, error)
		// 查询用户登录状态所有列表
		QueryMemberLoginStatusList(ctx context.Context, in *QueryMemberLoginStatusListReq, opts ...grpc.CallOption) (*QueryMemberLoginStatusListResp, error)
	}

	defaultService struct {
		cli zrpc.Client
	}
)

func NewService(cli zrpc.Client) Service {
	return &defaultService{
		cli: cli,
	}
}

// 添加用户登录状态
func (m *defaultService) AddMemberLoginStatus(ctx context.Context, in *AddMemberLoginStatusReq, opts ...grpc.CallOption) (*AddMemberLoginStatusResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.AddMemberLoginStatus(ctx, in, opts...)
}

// 更新用户登录状态
func (m *defaultService) UpdateMemberLogoutStatus(ctx context.Context, in *UpdateMemberLogoutStatusReq, opts ...grpc.CallOption) (*UpdateMemberLogoutStatusResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.UpdateMemberLogoutStatus(ctx, in, opts...)
}

// 查询用户登录状态详情
func (m *defaultService) QueryMemberLoginStatusDetail(ctx context.Context, in *QueryMemberLoginStatusDetailReq, opts ...grpc.CallOption) (*QueryMemberLoginStatusDetailResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.QueryMemberLoginStatusDetail(ctx, in, opts...)
}

// 查询用户登录状态分页列表
func (m *defaultService) QueryPageMemberLoginStatusList(ctx context.Context, in *QueryPageMemberLoginStatusListReq, opts ...grpc.CallOption) (*QueryPageMemberLoginStatusListResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.QueryPageMemberLoginStatusList(ctx, in, opts...)
}

// 查询用户登录状态所有列表
func (m *defaultService) QueryMemberLoginStatusList(ctx context.Context, in *QueryMemberLoginStatusListReq, opts ...grpc.CallOption) (*QueryMemberLoginStatusListResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.QueryMemberLoginStatusList(ctx, in, opts...)
}