package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"member/api/internal/logic"
	"member/api/internal/svc"
	"member/api/internal/types"
)

func QuerySysMemberListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySysMemberListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewQuerySysMemberListLogic(r.Context(), svcCtx)
		resp, err := l.QuerySysMemberList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
