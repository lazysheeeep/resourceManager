package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"resourceManager/rpc/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	DbClient *gorm.DB
	Mongodb  *mongo.Client
	Redis    *redis.Redis
}

func NewServiceContext(config config.Config) (*ServiceContext, error) {
	//Use the SetServerAPIOptions() method to set the stable A
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MongoDBConf.Protocol + "://localhost:" + config.MongoDBConf.Credentials).SetServerAPIOptions(serverAPI)

	//Create a new mongoClient and connect to the server
	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return &ServiceContext{}, err
	}

	//Create a new mysqlClient and connect to the server
	connRead := config.DatabaseConf.DbToString()
	connWrite := config.DatabaseConf.DbToString()
	db := config.DatabaseConf.NewDbClient(connRead, connWrite)

	return &ServiceContext{
		Config:   config,
		DbClient: db,
		Mongodb:  mongoClient,
		Redis:    redis.MustNewRedis(config.RedisConf),
	}, nil
}
