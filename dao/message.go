package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	Mid    string `json:"mid"`
	Uid    string `json:"uid"`
	Rid    string `json:"rid"`
	Data   string `json:"data"`
	Type   int    `json:"type"`
	Status int    `json:"status"`
	Ct     int64  `json:"ct"`
	Ut     int64  `json:"ut"`
}

func (Message) CollectionName() string {
	return "message"
}

func InsertMessage(message *Message) error {
	_, err := Mongo.Collection(Message{}.CollectionName()).InsertOne(context.Background(), message)
	return err
}

func FindMessageByRid(rid string, skip, limit *int64) (mes []*Message, err error) {
	mes = make([]*Message, 0)
	cur, err := Mongo.Collection(new(Message).CollectionName()).Find(context.Background(), bson.M{"rid": rid}, &options.FindOptions{Skip: skip, Limit: limit, Sort: bson.M{"ct": -1}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		message := new(Message)
		err = cur.Decode(message)
		if err != nil {
			return nil, err
		}
		mes = append(mes, message)
	}
	return
}
