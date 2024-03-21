package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"resourceManager/api/internal/logic/base"
	"resourceManager/api/internal/svc"
)

func InitDatbaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewInitDatbaseLogic(r.Context(), svcCtx)
		resp, err := l.InitDatbase()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
