package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Shop struct {
	ID      bson.ObjectId   `bson:"_id,omitempty"`
	Users   []string        `bson:"users"`
	Name    string          `bson:"name"`
	Phone   string          `bson:"phone"`
	Created time.Time       `bson:"created"`
	Config  ShopConfigs     `bson:"config"`
	Status  int             `bson:"status"`
	Theme   string          `bson:"theme"`
	Modules map[string]bool `bson:"modules"`
	Albums  ShopAlbum       `bson:"album"`
}

type ShopConfigs struct {
	Multilang   bool     `bson:"multilang"`
	UserDomain  bool     `bson:"userdomain"`
	Type        bool     `bons:"type"`
	Langs       []string `bson:"langs"`
	DefaultLang string   `bson:"defaultlang"`
}
type ShopLimit struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	ShopID string        `bson:"shopid"`
	Key    string        `bson:"key"`
	Value  int           `bson:"value"`
}

type ShopAlbum struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Slug    string        `bson:"slug"`
	Name    string        `bson:"name"`
	UserId  string        `bson:"userid"`
	ShopID  string        `bson:"shopid"`
	Created time.Time     `bson:"created"`
}
