package main

import (
	"context"
	"fmt"
	"im/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin123",
	}).ApplyURI("mongodb://0.0.0.0:27017"))
	if err != nil {
		log.Println("connection mongo error:", err)
	}
	db := client.Database("im")
	ub := new(model.UserBasic)
	err = db.Collection("user").FindOne(context.Background(), bson.D{}).Decode(ub)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ub ==> ", ub)
}
