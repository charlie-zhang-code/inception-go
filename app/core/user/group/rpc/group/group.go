// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: group.proto

package group

import (
	"context"

	"group/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddGroupReq                   = pb.AddGroupReq
	AddGroupResp                  = pb.AddGroupResp
	DeleteGroupReq                = pb.DeleteGroupReq
	DeleteGroupResp               = pb.DeleteGroupResp
	PageGroupListData             = pb.PageGroupListData
	QueryGroupDetailReq           = pb.QueryGroupDetailReq
	QueryGroupDetailResp          = pb.QueryGroupDetailResp
	QueryKeywordPageGroupListReq  = pb.QueryKeywordPageGroupListReq
	QueryKeywordPageGroupListResp = pb.QueryKeywordPageGroupListResp
	QueryPageGroupListReq         = pb.QueryPageGroupListReq
	QueryPageGroupListResp        = pb.QueryPageGroupListResp
	UpdateGroupReq                = pb.UpdateGroupReq
	UpdateGroupResp               = pb.UpdateGroupResp
	UpdateGroupStatusReq          = pb.UpdateGroupStatusReq
	UpdateGroupStatusResp         = pb.UpdateGroupStatusResp

	Group interface {
		// 添加用户组
		AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*AddGroupResp, error)
		// 删除用户组
		DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*DeleteGroupResp, error)
		// 更新用户组
		UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error)
		// 更新用户组状态
		UpdateGroupStatus(ctx context.Context, in *UpdateGroupStatusReq, opts ...grpc.CallOption) (*UpdateGroupStatusResp, error)
		// 查询用户组详情
		QueryGroupDetail(ctx context.Context, in *QueryGroupDetailReq, opts ...grpc.CallOption) (*QueryGroupDetailResp, error)
		// 查询用户组分页列表
		QueryPageGroupList(ctx context.Context, in *QueryPageGroupListReq, opts ...grpc.CallOption) (*QueryPageGroupListResp, error)
		// 查询用户组所有列表
		QueryKeywordPageGroupList(ctx context.Context, in *QueryKeywordPageGroupListReq, opts ...grpc.CallOption) (*QueryKeywordPageGroupListResp, error)
	}

	defaultGroup struct {
		cli zrpc.Client
	}
)

func NewGroup(cli zrpc.Client) Group {
	return &defaultGroup{
		cli: cli,
	}
}

// 添加用户组
func (m *defaultGroup) AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*AddGroupResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.AddGroup(ctx, in, opts...)
}

// 删除用户组
func (m *defaultGroup) DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*DeleteGroupResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.DeleteGroup(ctx, in, opts...)
}

// 更新用户组
func (m *defaultGroup) UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.UpdateGroup(ctx, in, opts...)
}

// 更新用户组状态
func (m *defaultGroup) UpdateGroupStatus(ctx context.Context, in *UpdateGroupStatusReq, opts ...grpc.CallOption) (*UpdateGroupStatusResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.UpdateGroupStatus(ctx, in, opts...)
}

// 查询用户组详情
func (m *defaultGroup) QueryGroupDetail(ctx context.Context, in *QueryGroupDetailReq, opts ...grpc.CallOption) (*QueryGroupDetailResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.QueryGroupDetail(ctx, in, opts...)
}

// 查询用户组分页列表
func (m *defaultGroup) QueryPageGroupList(ctx context.Context, in *QueryPageGroupListReq, opts ...grpc.CallOption) (*QueryPageGroupListResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.QueryPageGroupList(ctx, in, opts...)
}

// 查询用户组所有列表
func (m *defaultGroup) QueryKeywordPageGroupList(ctx context.Context, in *QueryKeywordPageGroupListReq, opts ...grpc.CallOption) (*QueryKeywordPageGroupListResp, error) {
	client := pb.NewGroupClient(m.cli.Conn())
	return client.QueryKeywordPageGroupList(ctx, in, opts...)
}
