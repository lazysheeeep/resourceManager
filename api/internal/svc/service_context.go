package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"resourceManager/api/internal/config"
	"resourceManager/rpc/coreclient"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware
	CoreRpc   coreclient.Core
	// McmsRpc mcmsclient.Mcms
	Redis  *redis.Redis
	Casbin *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithNewWatcher(c.DataBaseConf.Type, c.DataBaseConf.MysqlDSN(), c.RedisConf)

	return &ServiceContext{
		Config:  c,
		CoreRpc: coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:   rds,
		Casbin:  cbn,
	}
}
