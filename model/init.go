package model

import (
	"context"
	"fmt"
	"im/config"
	"im/help"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Database
var Redis *redis.Client

func Init(config config.Config) {
	Mongo = InitMongo(config.Mongo.Host, config.Mongo.Name, config.Mongo.Password, config.Mongo.Database, config.Mongo.Port)
	Redis = InitRedis(config.Redis.Host, config.Redis.Port)
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

func InitRedis(host string, port int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})
}
