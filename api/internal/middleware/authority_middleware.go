package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"resourceManager/api/internal/errors"
	"strings"
)

type AuthorityMiddleware struct {
	Csb *casbin.Enforcer
	Rds *redis.Redis
}

func NewAuthorityMiddleware(rds *redis.Redis, csb *casbin.Enforcer) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Csb: csb,
		Rds: rds,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//获取path
		obj := r.URL.Path
		//获取method 也就是act
		act := r.Method
		//获取用户身份
		roleId := r.Context().Value("roleId").(string)

		//检查token
		jwtResult, err := m.Rds.Get("token_" + StripBearerPrefixFromToken(r.Header.Get("Authorization")))
		if err != nil {
			logx.Errorf("redis error in jwt", logx.Field("detail:", err.Error()))
			httpx.Error(w, errors.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}

		if jwtResult == "1" {
			logx.Errorf("token is blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, errors.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		//检查权限
		result := batchCheck(m.Csb, roleId, act, obj)

		if result {
			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", r.Context().Value("userId").(string)),
				logx.Field("path", obj), logx.Field("method", act))
			next(w, r)
			return
		} else {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", roleId),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, errors.NewCodeError(7, "没有权限使用该接口"))
			return
		}

	}

}

func StripBearerPrefixFromToken(token string) string {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER" {
		return token[7:]
	}
	return token
}

func batchCheck(csb *casbin.Enforcer, roleId, act, obj string) bool {
	var checkReq [][]any
	for _, v := range strings.Split(roleId, ",") {
		checkReq = append(checkReq, []any{v, obj, act})
	}

	result, err := csb.BatchEnforce(checkReq)
	if err != nil {
		logx.Errorf("casbin enforce error", logx.Field("detail:", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
