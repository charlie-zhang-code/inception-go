// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: member.proto

package member

import (
	"context"

	"member/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CodeMessageResult           = pb.CodeMessageResult
	DeleteMemberIds             = pb.DeleteMemberIds
	MemberData                  = pb.MemberData
	QueryMemberDetailById       = pb.QueryMemberDetailById
	QueryMemberDetailByUsername = pb.QueryMemberDetailByUsername
	QueryPageMemberList         = pb.QueryPageMemberList
	QueryPageMemberListResult   = pb.QueryPageMemberListResult

	Member interface {
		// 添加用户
		CreateMember(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*CodeMessageResult, error)
		// 删除用户
		DeleteMember(ctx context.Context, in *DeleteMemberIds, opts ...grpc.CallOption) (*CodeMessageResult, error)
		// 更新用户
		UpdateMember(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*CodeMessageResult, error)
		// 查询用户详情（通过id）
		GetMemberDetailById(ctx context.Context, in *QueryMemberDetailById, opts ...grpc.CallOption) (*MemberData, error)
		// 查询用户详情（通过用户名）
		GetMemberDetailByUsername(ctx context.Context, in *QueryMemberDetailByUsername, opts ...grpc.CallOption) (*MemberData, error)
		// 查询用户分页列表
		GetKeywordPageMemberList(ctx context.Context, in *QueryPageMemberList, opts ...grpc.CallOption) (*QueryPageMemberListResult, error)
	}

	defaultMember struct {
		cli zrpc.Client
	}
)

func NewMember(cli zrpc.Client) Member {
	return &defaultMember{
		cli: cli,
	}
}

// 添加用户
func (m *defaultMember) CreateMember(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*CodeMessageResult, error) {
	client := pb.NewMemberClient(m.cli.Conn())
	return client.CreateMember(ctx, in, opts...)
}

// 删除用户
func (m *defaultMember) DeleteMember(ctx context.Context, in *DeleteMemberIds, opts ...grpc.CallOption) (*CodeMessageResult, error) {
	client := pb.NewMemberClient(m.cli.Conn())
	return client.DeleteMember(ctx, in, opts...)
}

// 更新用户
func (m *defaultMember) UpdateMember(ctx context.Context, in *MemberData, opts ...grpc.CallOption) (*CodeMessageResult, error) {
	client := pb.NewMemberClient(m.cli.Conn())
	return client.UpdateMember(ctx, in, opts...)
}

// 查询用户详情（通过id）
func (m *defaultMember) GetMemberDetailById(ctx context.Context, in *QueryMemberDetailById, opts ...grpc.CallOption) (*MemberData, error) {
	client := pb.NewMemberClient(m.cli.Conn())
	return client.GetMemberDetailById(ctx, in, opts...)
}

// 查询用户详情（通过用户名）
func (m *defaultMember) GetMemberDetailByUsername(ctx context.Context, in *QueryMemberDetailByUsername, opts ...grpc.CallOption) (*MemberData, error) {
	client := pb.NewMemberClient(m.cli.Conn())
	return client.GetMemberDetailByUsername(ctx, in, opts...)
}

// 查询用户分页列表
func (m *defaultMember) GetKeywordPageMemberList(ctx context.Context, in *QueryPageMemberList, opts ...grpc.CallOption) (*QueryPageMemberListResult, error) {
	client := pb.NewMemberClient(m.cli.Conn())
	return client.GetKeywordPageMemberList(ctx, in, opts...)
}