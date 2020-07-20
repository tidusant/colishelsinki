package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type UserSession struct {
	Session string
	UserID  string
	Action  string
	Params  string
	Shop    Shop
	UserIP  string
}
type User struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	User   string        `bson:"user"`
	Name   string        `bson:"name"`
	Email  string        `bson:"email"`
	Active int32         `bson:"active"`
	Group  string        `bson:"group"`
	Config UserConfig    `bson:"config"`
}

type UserConfig struct {
	Level     int `bson:"level"`
	MaxUpload int `bson:"maxupload"`
}

type UserLogin struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserId    bson.ObjectId `bson:"userid"`
	ShopId    string        `bson:"shopid"`
	Session   string        `bson:"session"`
	LastLogin time.Time     `bson:"last"`
	LoginIP   string        `bson:"ip"`
}
