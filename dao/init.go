package dao

import (
	"context"
	"fmt"
	"im/config"
	"im/help"
	"time"

	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Redis *redis.Client
var Mongo *mongo.Database

func Init() {
	Redis = InitRedis(config.Config.Redis.Host, config.Config.Redis.Port)
	Mongo = InitMongo(config.Config.Mongo.Host, config.Config.Mongo.Name, config.Config.Mongo.Password, config.Config.Mongo.Database, config.Config.Mongo.Port)
}

func InitRedis(host string, port int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})
}

func InitMongo(host, name, password, db string, port int) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: name,
		Password: password,
	}).ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		help.VczsLog("connection mongo error", err)
		return nil
	}
	return client.Database(db)
}
