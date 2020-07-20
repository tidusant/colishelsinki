package cuahang

import (
	c3mcommon "colis/common/common"

	"colis/models"
	"gopkg.in/mgo.v2/bson"
)

func GetCities() []models.City {
	col := db.C("geo_cities")
	var rs []models.City

	err := col.Find(bson.M{}).All(&rs)
	c3mcommon.CheckError("get cities", err)
	return rs
}
