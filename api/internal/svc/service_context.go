package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"resourceManager/api/internal/config"
	"resourceManager/api/internal/middleware"
	"resourceManager/rpc/coreclient"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	// McmsRpc mcmsclient.Mcms
	Redis   *redis.Redis
	Casbin  *casbin.Enforcer
	Captcha *base64Captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithNewWatcher(c.DataBaseConf.Type, c.DataBaseConf.MysqlDSN(), c.RedisConf)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(rds, cbn).Handle,
		CoreRpc:   coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:     rds,
		Casbin:    cbn,
		Captcha:   config.MustNewRedisCaptcha(c.Captcha, rds),
	}
}
