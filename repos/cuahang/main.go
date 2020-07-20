package cuahang

import (
	"os"

	c3mcommon "colis/common/common"
	"colis/common/log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	db *mgo.Database
)

func init() {
	log.Info("init repo cuahang...")
	strErr := ""
	db, strErr = c3mcommon.ConnectDB("chadmin")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
	log.Info("done")
}

//==============slug=================
func RemoveSlug(slug, shopid string) bool {
	col := db.C("addons_slugs")
	col.Remove(bson.M{"shopid": shopid, "slug": slug})
	return true
}
func CreateSlug(slug, shopid, object string) bool {
	col := db.C("addons_slugs")
	col.Insert(bson.M{"shopid": shopid, "slug": slug, "object": object})
	return true
}
func GetAllSlug(userid, shopid string) []string {
	col := db.C("addons_slugs")
	var rs []string
	cond := bson.M{"shopid": shopid}
	if userid != "594f665df54c58a2udfl54d3er" {
		cond["userid"] = userid
	}
	err := col.Find(cond).Select(bson.M{"slug": 1}).All(&rs)
	c3mcommon.CheckError("getallslug", err)
	return rs
}
