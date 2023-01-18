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

var Redis *redis.Client
var Mongo *mongo.Database

func Init(config *config.Config) error {
	redis := InitRedis(config.Redis.Host, config.Redis.Port)
	mongo, err := InitMongo(config.Mongo.Host, config.Mongo.Name, config.Mongo.Password, config.Mongo.Database, config.Mongo.Port)
	Redis = redis
	Mongo = mongo
	return err
}

func InitRedis(host string, port int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})
}

func InitMongo(host, name, password, db string, port int) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: name,
		Password: password,
	}).ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		help.VczsLog("connection mongo error", err)
		return nil, err
	}
	return client.Database(db), nil
}
