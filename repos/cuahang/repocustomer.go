package cuahang

import (
	c3mcommon "colis/common/common"
	"time"

	"colis/models"

	"gopkg.in/mgo.v2/bson"
)

func CountOrderByCus(phone, shopid string) int {
	col := db.C("addons_orders")
	rs := 0
	cond := bson.M{"shopid": shopid, "phone": phone}
	rs, err := col.Find(cond).Count()
	c3mcommon.CheckError("count order cus by phone", err)
	return rs
}
func GetAllCustomers(shopid string) []models.Customer {
	col := db.C("addons_customers")
	var rs []models.Customer
	cond := bson.M{"shopid": shopid}
	err := col.Find(cond).All(&rs)
	c3mcommon.CheckError("GetAllCustomers", err)
	return rs
}
func GetCusByPhone(phone, shopid string) models.Customer {
	col := db.C("addons_customers")
	var rs models.Customer
	cond := bson.M{"shopid": shopid, "phone": phone}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("get cus by phone", err)
	return rs
}
func GetCusByEmail(email, shopid string) models.Customer {
	col := db.C("addons_customers")
	var rs models.Customer
	cond := bson.M{"shopid": shopid, "email": email}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("get cus by email", err)
	return rs
}
func SaveCus(cus models.Customer) bool {
	col := db.C("addons_customers")
	cus.Modified = time.Now().UTC()
	if cus.ID == "" {
		cus.ID = bson.NewObjectId()
		cus.Created = cus.Modified
	}
	_, err := col.UpsertId(cus.ID, &cus)
	if c3mcommon.CheckError("save cus ", err) {
		return true
	}
	return false
}
