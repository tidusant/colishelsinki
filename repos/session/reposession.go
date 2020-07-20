package session

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	c3mcommon "colis/common/common"
	"colis/common/log"
	"colis/common/mystring"
	"colis/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	db *mgo.Database
)

func init() {
	fmt.Print("init repo session...")
	strErr := ""
	db, strErr = c3mcommon.ConnectDB("session")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
	fmt.Print("done\n")

}
//CreateSession: create session string and save into database
func CreateSession() string {
	sex := mystring.RandString(20)
	col := db.C("sessions")
	err := col.Insert(bson.M{"sex": sex, "created": time.Now().Unix(), "expired": time.Now().Unix() + 30*60})
	if c3mcommon.CheckError("Insert sessions", err) {
		return sex
	}
	return "-1"
}
func CheckSession(s string) bool {
	if s == "" {
		return false
	}
	col := db.C("sessions")
	var result models.Session
	err2 := col.Find(bson.M{"sex": s}).One(&result)

	if err2 != nil {
		log.Infof("Session not found sex '%s': %s\n", s, err2)
	} else {
		if result.Expired > time.Now().Unix() {
			//update session
			cond := bson.M{"_id": result.ID}
			change := bson.M{"$set": bson.M{"expired": time.Now().Unix() + 30*60}}
			col.Update(cond, change)
			return true
		} else {
			//remove session
			col.RemoveId(result.ID)
			log.Infof("Session expired: sex '%s'", s)
			return false
		}
	}
	return false
}

//CheckRequest: check request for anti ddos with request limit from env: REQUEST_LIMIT
func CheckRequest(uri, useragent, referrer, remoteAddress, requestType string) bool {
	col := db.C("requestUrls")
	//count same request url in 1 hour, if count>0 => already request => deny
	urlcount, _ := col.Find(bson.M{"uri": uri, "type": requestType, "created": bson.M{"$gt": int(time.Now().Unix()) - 1*3600}}).Count()
	if urlcount == 0 {
		//count url request of ip in 3 sec, if this ip request many time (requestlimit) => deny
		requestlimit,_:=strconv.Atoi(strings.Trim(os.Getenv("REQUEST_LIMIT")," "))
		if requestlimit==0{
			requestlimit=100
		}
		urlcount, err := col.Find(bson.M{"remoteAddress": remoteAddress, "created": bson.M{"$gt": int(time.Now().Unix()) - 3}}).Count()
		if urlcount < requestlimit {
			err = col.Insert(bson.M{"uri": uri, "created": int(time.Now().Unix()), "user-agent": useragent, "referer": referrer, "remoteAddress": remoteAddress, "type": requestType})
			c3mcommon.CheckError("checkRequest Insert", err)
			return true
		} else {
			log.Debugf("request ip limited %s", uri)
		}
	} else {
		log.Debugf("request limited %s", uri)
	}

	return false
}
