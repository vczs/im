package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type UserBasic struct {
	Uid      string `json:"uid"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Ct       int64  `json:"ct"`
	Ut       int64  `json:"ut"`
}

func (UserBasic) CollectionName() string {
	return "user"
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	user := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
