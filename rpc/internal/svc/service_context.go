package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"resourceManager/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *mongo.Client
	Redis  *redis.Redis
}

func NewServiceContext(config config.Config) (*ServiceContext, error) {
	//Use the SetServerAPIOptions() method to set the stable A
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MongoDBConf.Protocol + "://localhost:" + config.MongoDBConf.Credentials).SetServerAPIOptions(serverAPI)

	//Create a new client and connect to the server
	dbClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return &ServiceContext{}, err
	}

	return &ServiceContext{
		Config: config,
		DB:     dbClient,
		Redis:  redis.MustNewRedis(config.RedisConf),
	}, nil
}
