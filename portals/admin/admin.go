package main

import (
	c3mcommon "colis/common/common"
	"colis/common/log"
	"colis/common/mycrypto"
	pb "colis/grpcs/protoc"
	rpsex "colis/repos/session"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"time"

	"colis/models"
	//"io" repush
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {

}

var authport string
var authclusterIP string

//main function app run here
func main() {
	//expose port to outside of container
	var exposeport string
	exposeport = os.Getenv("EXPOSE_PORT")

	//get address of grpc auth from ENV
	authport = os.Getenv("AUTH_PORT")
	authclusterIP = os.Getenv("AUTH_CLUSTERIP")

	//other config:

	//default value for ENV
	if exposeport == "" {
		exposeport = "8081"
	}
	if authport == "" {
		authport = "32001"
	}

	//show info to console
	fmt.Println("auth address: " + authclusterIP + ":" + authport)
	fmt.Println("\n portal admin running with port " + exposeport)

	//start gin
	router := gin.Default()
	router.POST("/*name", postHandler)
	router.Run(":" + exposeport)

}

func postHandler(c *gin.Context) {
	strrt := ""
	requestDomain := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", requestDomain)
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers,access-control-allow-credentials")
	c.Header("Access-Control-Allow-Credentials", "true")

	//check request url, only one unique url per second
	if rpsex.CheckRequest(c.Request.URL.Path, c.Request.UserAgent(), c.Request.Referer(), c.Request.RemoteAddr, "POST") {
		rs := myRoute(c)
		b, _ := json.Marshal(rs)
		strrt = string(b)
	} else {
		log.Debugf("request denied")
	}

	if strrt == "" {
		strrt = c3mcommon.Fake64()
	} else {
		strrt = mycrypto.Encode(strrt, 8)
	}
	c.String(http.StatusOK, strrt)
}

func callgRPC(address string, rpcRequest pb.RPCRequest) models.RequestResult {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return c3mcommon.ReturnJsonMessage("-1", "service not run", "", "")
	}

	defer conn.Close()
	rpc := pb.NewGRPCServicesClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := rpc.Call(ctx, &rpcRequest)
	if err != nil {
		return c3mcommon.ReturnJsonMessage("-1", fmt.Sprintf("could not call service: %v", err.Error()), "", "")
	}
	return c3mcommon.ReturnJsonMessage("1", "", "", r.Data)
}

func myRoute(c *gin.Context) models.RequestResult {
	//get request name
	name := c.Param("name")
	name = name[1:] //remove  slash

	//get request data from Form
	data := c.PostForm("data")

	//get userip for check on 1 ip login
	userIP, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	//log.Debugf("decode name:%s", mycrypto.Decode(name))
	//decode request name and get array of args
	args := strings.Split(mycrypto.Decode(name), "|")

	//decode request data from Form and get array of args
	datargs := strings.Split(mycrypto.Decode(data), "|")
	requestAction := datargs[0]
	requestParams := ""
	if len(datargs) > 1 {
		requestParams = datargs[1]
	}

	//get rpc call name from first arg
	RPCname := args[0]

	if RPCname == "CreateSex" {
		//create session string and save it into db
		data = rpsex.CreateSession()

		return c3mcommon.ReturnJsonMessage("1", "", "", data)
	}

	session := ""
	if len(args) > 1 {
		session = args[1]
	}

	//get session from other server's call
	if session == "" && datargs[0] == "test" && len(datargs) > 1 {
		session = mycrypto.Decode(datargs[1])
	}

	reply := c3mcommon.ReturnJsonMessage("0", "unknown error", "", "")

	//check session
	if !rpsex.CheckSession(session) {
		return c3mcommon.ReturnJsonMessage("-2", "session not found", "", "")
	}
	if RPCname == "aut" {
		return callgRPC(authclusterIP+":"+authport, pb.RPCRequest{AppName: "admin-portal", Action: requestAction, Params: requestParams, Session: session, UserIP: userIP})
	}

	//always check login if RPCname not aut and create session
	authreply := callgRPC(authclusterIP+":"+authport, pb.RPCRequest{AppName: "admin-portal", Action: "auth", Params: requestParams, Session: session, UserIP: userIP})
	if authreply.Status != "1" {
		return authreply
	}
	//get logininfo: from check login in format: userid[+]shopid
	logininfo := strings.Split(reply.Data, "[+]")
	UserId := logininfo[0]
	ShopId := logininfo[1]
	//begin gRPC call
	return callgRPC(authclusterIP+":"+authport, pb.RPCRequest{AppName: "admin-portal", Action: requestAction, Params: requestParams, Session: session, UserIP: userIP, UserID: UserId, ShopID: ShopId})

}
