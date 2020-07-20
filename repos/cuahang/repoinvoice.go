package cuahang

import (
	c3mcommon "colis/common/common"

	"colis/models"
	//	"c3m/log"

	//"strings"

	"gopkg.in/mgo.v2/bson"
)

func SaveInvoice(invc models.Invoice) models.Invoice {

	col := db.C("addons_invoice")

	if invc.ID == "" {
		invc.ID = bson.NewObjectId()
	}
	col.UpsertId(invc.ID, &invc)
	return invc
}

func GetInvoices(shopid string, imp bool) []models.Invoice {

	col := db.C("addons_invoice")

	var rs []models.Invoice
	err := col.Find(bson.M{"shopid": shopid, "import": imp}).Sort("created").All(&rs)
	c3mcommon.CheckError("GetInvoices", err)
	return rs
}
func GetInvcById(shopid, invcid string) models.Invoice {

	col := db.C("addons_invoice")
	var rs models.Invoice
	err := col.Find(bson.M{"shopid": shopid, "_id": bson.ObjectIdHex(invcid)}).One(&rs)
	c3mcommon.CheckError("GetInvcById", err)

	return rs
}
func RemoveInvcById(shopid, invcid string) bool {

	col := db.C("addons_invoice")

	err := col.RemoveId(bson.ObjectIdHex(invcid))
	c3mcommon.CheckError("RemoveInvcById", err)
	return true
}
