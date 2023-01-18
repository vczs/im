package dao

import "context"

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
