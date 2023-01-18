package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	Uid    string `json:"uid"`
	Rid    string `json:"rid"`
	Type   int    `json:"type"`
	Status int    `json:"status"`
	Ct     int64  `json:"ct"`
	Ut     int64  `json:"ut"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByUidRid(uid, rid string) (num int64, err error) {
	return Mongo.Collection(UserRoom{}.CollectionName()).CountDocuments(context.Background(), bson.D{{"uid", uid}, {"rid", rid}})
}

func FindUserRoomByRid(rid string) (urs []*UserRoom, err error) {
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"rid", rid}})
	if err != nil {
		return nil, err
	}
	urs = make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}
