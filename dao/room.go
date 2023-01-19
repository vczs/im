package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Room struct {
	Rid      string `json:"rid"`
	Uid      string `json:"uid"`
	Number   int    `json:"number"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Status   int    `json:"status"`
	Ct       int64  `json:"ct"`
	Ut       int64  `json:"ut"`
}

func (Room) CollectionName() string {
	return "room"
}

func InsertRoom(room *Room) error {
	_, err := Mongo.Collection(Room{}.CollectionName()).InsertOne(context.Background(), room)
	return err
}

func DeleteRoomByRid(rid string) error {
	_, err := Mongo.Collection(new(Room).CollectionName()).UpdateMany(context.Background(), bson.M{"rid": rid}, bson.M{"$set": bson.M{"status": -1, "ut": time.Now().Unix()}})
	return err
}
