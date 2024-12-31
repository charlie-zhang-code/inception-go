// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"member/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/addSysUser",
				Handler: AddSysUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/deleteSysUser",
				Handler: DeleteSysUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/queryPageSysUserList",
				Handler: QueryPageSysUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/querySysUserDetail",
				Handler: QuerySysUserDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/querySysUserList",
				Handler: QuerySysUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateSysUser",
				Handler: UpdateSysUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateSysUserStatus",
				Handler: UpdateSysUserStatusHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/user/member"),
	)
}
