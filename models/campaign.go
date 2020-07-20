package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Campaign struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	ShopId         string        `bson:"shopid"`
	Name           string        `bson:"name"`
	Description    string        `bson:"description"`
	Bugget         int           `bson:"bugget"`
	UserId         string        `bson:"userid"`
	Created        time.Time     `bson:"created"`
	Modified       time.Time     `bson:"modified"`
	Start          time.Time     `bson:"start"`
	End            time.Time     `bson:"end"`
	Noo            int
	Total          int
	Base           int
	PartnerShipFee int
	ShipFee        int
	StatusDetail   map[string]CampaignStatusDetail
}

type CampaignStatusDetail struct {
	Total          int
	Base           int
	Noo            int
	PartnerShipFee int
	ShipFee        int
}
