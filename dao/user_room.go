package dao

import (
	"context"
	"im/define"
	"time"

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

func InsertManyUserRoom(urs []interface{}) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).InsertMany(context.Background(), urs)
	return err
}

func GetUserRoomCountByUidRid(uid, rid string) (num int64, err error) {
	return Mongo.Collection(UserRoom{}.CollectionName()).CountDocuments(context.Background(), bson.M{"uid": uid, "rid": rid})
}

func FindUserRoomByRid(rid string) (urs []*UserRoom, err error) {
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.M{"rid": rid})
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

// 判断是否为好友
func IsFriend(u1, u2 string) (isFriend bool, rid string) {
	isFriend = false
	ur := new(UserRoom)
	cur, err := Mongo.Collection(new(UserRoom).CollectionName()).Find(context.Background(), bson.M{"uid": u1, "type": define.UserRoomTypeAlone, "status": 1})
	if err != nil {
		return
	}
	rs := make([]string, 0)
	for cur.Next(context.Background()) {
		ur := new(UserRoom)
		err = cur.Decode(ur)
		if err != nil {
			return
		}
		rs = append(rs, ur.Rid)
	}
	err = Mongo.Collection(new(UserRoom).CollectionName()).FindOne(context.Background(), bson.M{"uid": u2, "type": define.UserRoomTypeAlone, "status": 1, "rid": bson.M{"$in": rs}}).Decode(ur)
	if err != nil {
		return
	}
	if ur.Rid != "" {
		rid = ur.Rid
		isFriend = true
	}
	return
}

func DeleteUserRoomByRid(rid string) error {
	_, err := Mongo.Collection(new(UserRoom).CollectionName()).UpdateMany(context.Background(), bson.M{"rid": rid}, bson.M{"$set": bson.M{"status": -1, "ut": time.Now().Unix()}})
	return err
}
