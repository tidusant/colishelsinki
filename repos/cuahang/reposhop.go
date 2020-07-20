package cuahang

import (
	c3mcommon "colis/common/common"
	"colis/common/log"
	"colis/models"

	"gopkg.in/mgo.v2/bson"
)

/*for dashboard
=============================================================================
*/
func UpdateTheme(userid, shopid, code string) string {
	col := db.C("addons_shops")

	shop := GetShopById(userid, shopid)
	if shop.Name != "" {
		log.Debugf("update login info:shopid %s", shop.ID.Hex())
		change := bson.M{"$set": bson.M{"theme": code}}
		err := col.UpdateId(shop.ID, change)
		c3mcommon.CheckError("update theme", err)
	}
	return shop.Name
}
func LoadShopById(session, userid, shopid string) models.Shop {
	col := db.C("addons_userlogin")
	if shopid == "" {
		//get first shop
		shopid = GetShopDefault(userid)
	}
	shop := GetShopById(userid, shopid)
	if shop.Name != "" {
		log.Debugf("update login info:shopid %s", shop.ID.Hex())
		cond := bson.M{"session": session, "userid": bson.ObjectIdHex(userid)}
		change := bson.M{"$set": bson.M{"shopid": shop.ID.Hex(), "shoplang": shop.Config.DefaultLang}}
		col.Update(cond, change)
	}
	return shop
}
func GetShopDefault(userid string) string {
	col := db.C("addons_shops")
	var result models.Shop

	col.Find(bson.M{"users.id": userid}).One(&result)
	if result.Name != "" {
		return result.ID.Hex()
	}

	//pipeline := []bson.M{{"$match": bson.M{"name": "abc"}}}
	//col.Pipe(pipeline).All(&result)
	//	for {
	//		if iter.Next(&result) {
	//			log.Printf("result %v", result)
	//			return result.ID.Hex()
	//		} else {
	//			return ""
	//		}
	//	}

	//	if len(result) > 0 {
	//		return result[0].ID.Hex()
	//	}
	return ""
}
func GetShopById(userid, shopid string) models.Shop {
	coluser := db.C("addons_shops")
	var shop models.Shop
	if shopid == "" {
		return shop
	}
	cond := bson.M{"_id": bson.ObjectIdHex(shopid)}

	coluser.Find(cond).One(&shop)
	return shop
}
func GetOtherShopById(userid, shopid string) []models.Shop {
	coluser := db.C("addons_shops")
	var shops []models.Shop
	if shopid == "" {
		return shops
	}
	cond := bson.M{"_id": bson.M{"$ne": bson.ObjectIdHex(shopid)}}
	cond["users.id"] = userid

	coluser.Find(cond).All(&shops)
	return shops
}
func GetDemoShop() models.Shop {
	coluser := db.C("addons_shops")
	var shop models.Shop
	coluser.Find(bson.M{"name": "demo"}).One(&shop)
	return shop
}

// func SaveCat(userid, shopid string, cat models.ProdCat) string {

// 	shop := GetShopById(userid, shopid)
// 	newcat := false
// 	if cat.Code == "" {
// 		newcat = true
// 	}
// 	//get all cats
// 	cats := GetAllCats(userid, shopid)
// 	var oldcat models.ProdCat
// 	//check max cat limited
// 	if shop.Config.MaxCat <= len(cats) && newcat {
// 		return "-1"
// 	}
// 	//get array of album slug
// 	catslugs := map[string]string{}
// 	catcodes := map[string]string{}
// 	for _, c := range cats {
// 		catcodes[c.Code] = c.Code
// 		for _, ci := range c.Langs {
// 			catslugs[ci.Slug] = ci.Slug
// 		}
// 		if newcat && c.Code == cat.Code {
// 			oldcat = c
// 		}
// 	}

// 	for lang, _ := range cat.Langs {
// 		if cat.Langs[lang].Name == "" {
// 			delete(cat.Langs, lang)
// 			continue
// 		}
// 		//newslug
// 		i := 1
// 		newslug := inflect.Parameterize(cat.Langs[lang].Name)
// 		cat.Langs[lang].Slug = newslug
// 		//check slug duplicate
// 		for {
// 			if _, ok := catslugs[cat.Langs[lang].Slug]; ok {
// 				cat.Langs[lang].Slug = newslug + strconv.Itoa(i)
// 				i++
// 			} else {
// 				catslugs[cat.Langs[lang].Slug] = cat.Langs[lang].Slug
// 				break
// 			}
// 		}
// 	}

// 	//check code duplicate
// 	if newcat {
// 		//insert new
// 		newcode := ""
// 		for {
// 			newcode = mystring.RandString(3)
// 			if _, ok := catcodes[newcode]; !ok {
// 				break
// 			}
// 		}
// 		cat.Code = newcode
// 		cat.UserId = userid
// 		cat.Created = time.Now().UTC().Add(time.Hour + 7)
// 	} else {
// 		//update
// 		oldcat.Langs = cat.Langs
// 		cat = oldcat
// 	}

// 	UpdateCat(shop)
// 	return cat.Code
// }

//func SaveCat(userid, shopid, code string, catinfo models.ShopCatInfo) string {

//	//slug
//	rt := "0"
//	i := 1
//	slug := inflect.Parameterize(catinfo.Name)
//	catinfo.Slug = slug
//	shop := GetShopById(userid, shopid)

//	//get array of album slug
//	catslugs := map[string]string{}
//	for _, c := range shop.ShopCats {
//		for _, ci := range c.Langs {
//			if ci.Slug != catinfo.Slug {
//				catslugs[ci.Slug] = ci.Slug
//			}
//		}

//	}

//	for {
//		if _, ok := catslugs[catinfo.Slug]; ok {
//			catinfo.Slug = slug + strconv.Itoa(i)
//			i++
//		} else {
//			break
//		}
//	}

//	for i, _ := range shop.ShopCats {
//		if shop.ShopCats[i].Code == code && shop.ShopCats[i].UserId == userid {
//			isnewlang := true
//			for j, _ := range shop.ShopCats[i].Langs {
//				if shop.ShopCats[i].Langs[j].Lang == catinfo.Lang {
//					//shop.ShopCats[i].Langs[j] = catinfo
//					isnewlang = false
//					break
//				}
//			}
//			if isnewlang {
//				//shop.ShopCats[i].Langs = append(shop.ShopCats[i].Langs, catinfo)

//			}
//			rt = "1"
//			break
//		}
//	}
//	UpdateCat(shop)
//	return rt

//}
func UpdateAlbum(shop models.Shop) models.Shop {
	coluser := db.C("addons_shops")

	cond := bson.M{"_id": shop.ID}
	change := bson.M{"$set": bson.M{"albums": shop.Albums}}

	coluser.Update(cond, change)
	return shop
}
