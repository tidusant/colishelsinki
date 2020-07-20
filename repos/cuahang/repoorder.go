package cuahang

import (
	c3mcommon "colis/common/common"
	"colis/common/log"
	"colis/common/mystring"
	"colis/models"

	"time"

	"gopkg.in/mgo.v2/bson"
)

func GetOrderByID(orderid, shopid string) models.Order {
	col := db.C("addons_orders")
	var rs models.Order
	cond := bson.M{"shopid": shopid, "_id": bson.ObjectIdHex(orderid)}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("GetOrdersByID", err)
	return rs
}

func GetOrdersByStatus(shopid, status string, page int, pagesize int, searchterm string) []models.Order {
	col := db.C("addons_orders")
	var rs []models.Order

	// pipeline := []bson.M{
	// 	{"$match": bson.M{"shopid": shopid, "status": status}},
	// 	{"$sort": bson.D{
	// 		{"_id", -1}}, //1: Ascending, -1: Descending
	// 	},

	// 	{"$skip": (page - 1) * pagesize}, //1: Ascending, -1: Descending

	// 	{"$limit": pagesize}, //1: Ascending, -1: Descending

	// }
	// pipe := col.Pipe(pipeline)
	// //err := pipe.AllowDiskUse().All(&result) //allow disk use
	// err := pipe.All(&rs)

	cond := bson.M{"shopid": shopid}
	if status != "all" {
		cond["status"] = status
	}
	if searchterm != "" {
		//searchtermslug := strings.Replace(searchterm, " ", "-", -1)
		searchtermslug := mystring.ParameterizeJoin(searchterm, " ")
		log.Debugf("searchteram slug: $s", searchtermslug)
		//searchtermslug = strings.Replace(searchtermslug, "-", " ", -1)
		//log.Debugf("searchteram slug: $s", searchtermslug)
		cond["$or"] = []bson.M{
			bson.M{"searchindex": bson.M{"$regex": bson.RegEx{searchtermslug, "si"}}},
			// bson.M{"email": bson.M{"$regex": bson.RegEx{searchterm, "si"}}},
			// bson.M{"name": bson.M{"$regex": bson.RegEx{searchterm, "si"}}},
			// bson.M{"name": bson.M{"$regex": bson.RegEx{searchtermslug, "si"}}},
			// bson.M{"address": bson.M{"$regex": bson.RegEx{searchterm, "si"}}},
			// bson.M{"address": bson.M{"$regex": bson.RegEx{searchtermslug, "si"}}},
			// bson.M{"note": bson.M{"$regex": bson.RegEx{searchterm, "si"}}},
			// bson.M{"note": bson.M{"$regex": bson.RegEx{searchtermslug, "si"}}},
		}
	}

	var err error
	if page == 0 {
		err = col.Find(cond).Sort("-_id").All(&rs)
	} else {
		err = col.Find(cond).Sort("-_id").Skip((page - 1) * pagesize).Limit(pagesize).All(&rs)
	}
	c3mcommon.CheckError("GetOrdersByStatus", err)
	return rs
}
func CountOrdersByStatus(shopid, status, searchterm string) int {
	col := db.C("addons_orders")

	cond := bson.M{"shopid": shopid}
	if status != "all" {
		cond["status"] = status
	}
	if searchterm != "" {
		searchtermslug := mystring.ParameterizeJoin(searchterm, " ")
		log.Debugf("searchteram slug: $s", searchtermslug)
		//searchtermslug = strings.Replace(searchtermslug, "-", " ", -1)
		//log.Debugf("searchteram slug: $s", searchtermslug)
		cond["$or"] = []bson.M{
			bson.M{"searchindex": bson.M{"$regex": bson.RegEx{searchtermslug, "si"}}},
		}
	}

	count, err := col.Find(cond).Count()
	log.Debugf("count search: %v", count)
	c3mcommon.CheckError("CountOrdersByStatus", err)
	return count
}
func GetOrdersByCampaign(camp models.Campaign) []models.Order {
	col := db.C("addons_orders")
	var rs []models.Order
	cond := bson.M{"shopid": camp.ShopId}
	err := col.Find(cond).All(&rs)
	c3mcommon.CheckError("GetOrdersByCamp", err)
	return rs
}

func GetDefaultOrderStatus(shopid string) models.OrderStatus {
	col := db.C("addons_order_status")
	var rs models.OrderStatus

	cond := bson.M{"shopid": shopid, "default": true}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("GetDefaultOrderStatus", err)
	return rs
}

func UpdateOrderStatus(shopid, status string, orderid string) {
	col := db.C("addons_orders")
	//var arrIdObj []bson.ObjectId
	// for _, v := range orderid {
	// 	arrIdObj = append(arrIdObj, bson.ObjectIdHex(v))
	// }
	// cond := bson.M{"_id": bson.M{"$in": arrIdObj}, "shopid": shopid}
	cond := bson.M{"_id": bson.ObjectIdHex(orderid), "shopid": shopid}
	change := bson.M{"status": status}
	stats := GetAllOrderStatus(shopid)
	mapstat := make(map[string]models.OrderStatus)
	for _, v := range stats {
		mapstat[v.ID.Hex()] = v
	}

	if mapstat[status].Finish {
		change["whookupdate"] = time.Now().Unix()
	}
	log.Debugf("udpate order cond:%v", cond)
	err := col.Update(cond, bson.M{"$set": change})
	c3mcommon.CheckError("Update order status", err)
	return
}

func SaveOrder(order models.Order) models.Order {
	col := db.C("addons_orders")
	if order.ID == "" {
		order.ID = bson.NewObjectId()
		order.Created = time.Now().Unix()

	}

	order.Modified = time.Now().Unix()
	col.UpsertId(order.ID, order)
	return order
}

func GetCountOrderByStatus(stat models.OrderStatus) int {
	col := db.C("addons_orders")

	cond := bson.M{"shopid": stat.ShopId, "status": stat.ID.Hex()}
	n, err := col.Find(cond).Count()
	c3mcommon.CheckError("GetstatusByID", err)
	return n
}

//====================== whook

//===============status function
func UpdateOrderStatusByShipmentCode(shipmentCode, statusid, shopid string) {
	col := db.C("addons_orders")
	cond := bson.M{"shopid": shopid, "shipmentcode": shipmentCode}
	change := bson.M{"$set": bson.M{"status": statusid, "whookupdate": time.Now().Unix()}}

	err := col.Update(cond, change)
	c3mcommon.CheckError("UpdateOrderStatusByShipmentCode", err)
}
func GetOrderByShipmentCode(shipmentCode, shopid string) models.Order {
	col := db.C("addons_orders")
	var rs models.Order
	cond := bson.M{"shopid": shopid, "shipmentcode": shipmentCode}
	err := col.Find(cond).One(&rs)

	c3mcommon.CheckError("GetOrderByShipmentCode", err)
	return rs
}
func GetStatusByPartnerStatus(partnercode, shopid, partnerstatus string) models.OrderStatus {
	col := db.C("addons_order_status")
	var rs models.OrderStatus

	cond := bson.M{"shopid": shopid, "partnerstatus." + partnercode: partnerstatus}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("GetStatusByPartnerStatus", err)
	return rs
}
func GetAllOrderStatus(shopid string) []models.OrderStatus {
	col := db.C("addons_order_status")
	var rs []models.OrderStatus
	cond := bson.M{"shopid": shopid}
	err := col.Find(cond).All(&rs)
	c3mcommon.CheckError("GetAllOrderStatus", err)
	return rs
}
func SaveOrderStatus(status models.OrderStatus) models.OrderStatus {
	col := db.C("addons_order_status")
	if status.ID == "" {
		status.ID = bson.NewObjectId()
		status.Created = time.Now().UTC()
	}

	status.Modified = status.Created
	col.UpsertId(status.ID, status)
	return status
}

func GetStatusByID(statusid, shopid string) models.OrderStatus {
	col := db.C("addons_order_status")
	var rs models.OrderStatus
	cond := bson.M{"shopid": shopid, "_id": bson.ObjectIdHex(statusid)}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("GetstatusByID", err)
	return rs
}

func DeleteOrderStatus(stat models.OrderStatus) bool {
	col := db.C("addons_order_status")

	cond := bson.M{"shopid": stat.ShopId, "_id": stat.ID}
	err := col.Remove(cond)
	return c3mcommon.CheckError("GetstatusByID", err)

}

func UnSetStatusDefault(shopid string) {
	col := db.C("addons_order_status")

	cond := bson.M{"shopid": shopid, "default": true}
	change := bson.M{"$set": bson.M{"default": false}}
	err := col.Update(cond, change)
	c3mcommon.CheckError("GetstatusByDefault", err)

}
