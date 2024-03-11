package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"resourceManager/api/internal/config"
	"resourceManager/rpc/coreclient"
)

type ServiceContext struct {
	Config config.Config
	// Authority rest.Middleware
	CoreRpc coreclient.Core
	// McmsRpc mcmsclient.Mcms
	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:  c,
		CoreRpc: coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Redis:   rds,
	}
}
