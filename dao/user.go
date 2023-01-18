package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
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

func (User) CollectionName() string {
	return "user"
}

// 通过uid获取用户
func GetUserByUid(uid string) (user *User, err error) {
	user = new(User)
	err = Mongo.Collection(User{}.CollectionName()).FindOne(context.Background(), bson.D{{"uid", uid}}).Decode(user)
	return
}

// 通过account、password获取用户
func GetUserByAccountPassword(account, password string) (user *User, err error) {
	user = new(User)
	err = Mongo.Collection(User{}.CollectionName()).FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).Decode(user)
	return
}

// 通过email查询用户数量
func GetUserCountByEmail(email string) (num int64, err error) {
	return Mongo.Collection(User{}.CollectionName()).CountDocuments(context.Background(), bson.D{{"email", email}})
}
