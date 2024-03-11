package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"resourceManager/api/internal/logic/user"
	"resourceManager/api/internal/svc"
	"resourceManager/api/internal/types"
)

func GetUserByUsernameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UsernameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserByUsernameLogic(r.Context(), svcCtx)
		resp, err := l.GetUserByUsername(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
