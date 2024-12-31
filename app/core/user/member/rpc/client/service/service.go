// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: main.proto

package service

import (
	"context"

	"member/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddSysUserReq            = pb.AddSysUserReq
	AddSysUserResp           = pb.AddSysUserResp
	DeleteSysUserReq         = pb.DeleteSysUserReq
	DeleteSysUserResp        = pb.DeleteSysUserResp
	PageSysUserListData      = pb.PageSysUserListData
	QueryPageSysUserListReq  = pb.QueryPageSysUserListReq
	QueryPageSysUserListResp = pb.QueryPageSysUserListResp
	QuerySysUserDetailReq    = pb.QuerySysUserDetailReq
	QuerySysUserDetailResp   = pb.QuerySysUserDetailResp
	QuerySysUserListReq      = pb.QuerySysUserListReq
	QuerySysUserListResp     = pb.QuerySysUserListResp
	SysUserListData          = pb.SysUserListData
	UpdateSysUserReq         = pb.UpdateSysUserReq
	UpdateSysUserResp        = pb.UpdateSysUserResp
	UpdateSysUserStatusReq   = pb.UpdateSysUserStatusReq
	UpdateSysUserStatusResp  = pb.UpdateSysUserStatusResp

	Service interface {
		// 添加用户
		AddSysUser(ctx context.Context, in *AddSysUserReq, opts ...grpc.CallOption) (*AddSysUserResp, error)
		// 删除用户
		DeleteSysUser(ctx context.Context, in *DeleteSysUserReq, opts ...grpc.CallOption) (*DeleteSysUserResp, error)
		// 更新用户
		UpdateSysUser(ctx context.Context, in *UpdateSysUserReq, opts ...grpc.CallOption) (*UpdateSysUserResp, error)
		// 更新用户状态
		UpdateSysUserStatus(ctx context.Context, in *UpdateSysUserStatusReq, opts ...grpc.CallOption) (*UpdateSysUserStatusResp, error)
		// 查询用户详情
		QuerySysUserDetail(ctx context.Context, in *QuerySysUserDetailReq, opts ...grpc.CallOption) (*QuerySysUserDetailResp, error)
		// 查询用户分页列表
		QueryPageSysUserList(ctx context.Context, in *QueryPageSysUserListReq, opts ...grpc.CallOption) (*QueryPageSysUserListResp, error)
		// 查询用户所有列表
		QuerySysUserList(ctx context.Context, in *QuerySysUserListReq, opts ...grpc.CallOption) (*QuerySysUserListResp, error)
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

// 添加用户
func (m *defaultService) AddSysUser(ctx context.Context, in *AddSysUserReq, opts ...grpc.CallOption) (*AddSysUserResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.AddSysUser(ctx, in, opts...)
}

// 删除用户
func (m *defaultService) DeleteSysUser(ctx context.Context, in *DeleteSysUserReq, opts ...grpc.CallOption) (*DeleteSysUserResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.DeleteSysUser(ctx, in, opts...)
}

// 更新用户
func (m *defaultService) UpdateSysUser(ctx context.Context, in *UpdateSysUserReq, opts ...grpc.CallOption) (*UpdateSysUserResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.UpdateSysUser(ctx, in, opts...)
}

// 更新用户状态
func (m *defaultService) UpdateSysUserStatus(ctx context.Context, in *UpdateSysUserStatusReq, opts ...grpc.CallOption) (*UpdateSysUserStatusResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.UpdateSysUserStatus(ctx, in, opts...)
}

// 查询用户详情
func (m *defaultService) QuerySysUserDetail(ctx context.Context, in *QuerySysUserDetailReq, opts ...grpc.CallOption) (*QuerySysUserDetailResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.QuerySysUserDetail(ctx, in, opts...)
}

// 查询用户分页列表
func (m *defaultService) QueryPageSysUserList(ctx context.Context, in *QueryPageSysUserListReq, opts ...grpc.CallOption) (*QueryPageSysUserListResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.QueryPageSysUserList(ctx, in, opts...)
}

// 查询用户所有列表
func (m *defaultService) QuerySysUserList(ctx context.Context, in *QuerySysUserListReq, opts ...grpc.CallOption) (*QuerySysUserListResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.QuerySysUserList(ctx, in, opts...)
}