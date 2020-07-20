package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"strings"

	c3mcommon "colis/common/common"
	"colis/common/log"
	pb "colis/grpcs/protoc"
	"colis/models"
	rpch "colis/repos/cuahang"
	"context"
	"google.golang.org/grpc"
)

const (
	name string ="auth"
 ver string = "1"
)

type service struct {
	pb.UnimplementedGRPCServicesServer
}
func (s *service) Call(ctx context.Context, in *pb.RPCRequest) (*pb.RPCResponse, error) {
	rs:=&pb.RPCResponse{Data: "Hello " + in.GetAppName(),RPCName: name, Version: ver}
	result := models.RequestResult{}
	//get input data into user session
	var usex models.UserSession
	usex.Session = in.Session
	usex.Action = in.Action
	usex.UserID=in.UserID
	usex.UserIP=in.UserIP
	usex.Params = in.Params

	if usex.Action == "l" {
		result = login(usex)
	} else if usex.Action == "lo" {
		result = logout(usex.Session)
	} else if usex.Action == "test" {
		result = test(usex)
	} else if usex.Action == "aut" {
		result=auth(usex)
	}else{
		return rs, nil
	}
	//convert RequestResult into json
	b,_:=json.Marshal(result)
	rs.Data=string(b)
	return rs, nil
}
//test: to authenticate user already login, for frontend, return session only
func test(usex models.UserSession) models.RequestResult {
	if rpch.GetLogin(usex.Session) != "" {
		return c3mcommon.ReturnJsonMessage("1", "", "user logged in", `{"sex":"`+usex.Session+`"}`)
	}
	return c3mcommon.ReturnJsonMessage("0", "user not logged in", "", `{"sex":"`+usex.Session+`"}`)
}
//auth: to authenticate user already login, for portal, return userid[+]shopid
func auth(usex models.UserSession) models.RequestResult {
	logininfo := rpch.GetLogin(usex.Session)
	if logininfo == "" {
		return c3mcommon.ReturnJsonMessage("0", "user not logged in", "", "")
	} else {
		return c3mcommon.ReturnJsonMessage("1", "", "user logged in", logininfo)
	}
}
//login user and update Session and IP in user_login.
func login(usex models.UserSession) models.RequestResult {
	args := strings.Split(usex.Params, ",")
	if len(args) < 2 {
		return c3mcommon.ReturnJsonMessage("0", "empty username or pass", "", "")
	}
	user := args[0]
	pass := args[1]
	name := rpch.Login(user, pass, usex.Session, usex.UserIP)
	if name != "" {
		return c3mcommon.ReturnJsonMessage("1", "", "login success", name)
	}
	return c3mcommon.ReturnJsonMessage("0", "login fail", "", "")

}

func logout(session string) models.RequestResult {
	rpch.Logout(session)

	return c3mcommon.ReturnJsonMessage("1", "", "login success", "")

}
func main() {
	//default port for service
	var port string
	port=os.Getenv("PORT")
	if port==""{
		port="8901"
	}
	//open service and listen
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Printf("listening on %s\n",port)
	pb.RegisterGRPCServicesServer(s, &service{})
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve : %v", err)
	}
	fmt.Print("exit")

}
