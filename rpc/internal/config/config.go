package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf DatabaseConf
	MongoDBConf  MongoDBConf
	RedisConf    redis.RedisConf
	QiniuConf    QiNiuConf
}
