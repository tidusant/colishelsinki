package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	ShopId   string        `bson:"shopid"`
	Name     string        `bson:"name"`
	Phone    string        `bson:"phone"`
	City     string        `bson:"city"`
	District string        `bson:"district"`
	Ward     string        `bson:"ward"`
	Address  string        `bson:"address"`
	Note     string        `bson:"note"`
	Email    string        `bson:"email"`
	Created  time.Time     `bson:"created"`
	Modified time.Time     `bson:"modified"`
}
