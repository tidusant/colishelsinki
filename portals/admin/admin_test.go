package main

import (
	"bytes"
	c3mcommon "colis/common/common"
	"colis/common/mycrypto"
	"colis/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var testsession string

func decodeResponse(requeststring string, data string) (rs models.RequestResult, errstr string) {
	//encode data
	requeststring = mycrypto.EncDat2(requeststring)
	data = "data=" + mycrypto.EncDat2(data)

	//add body into request
	body := bytes.NewReader([]byte(data))
	req, err := http.NewRequest(http.MethodPost, "/"+requeststring, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		errstr = fmt.Sprintf("Couldn't create request: %v\n", err)
		return
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		errstr = fmt.Sprintf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
		return

	}

	//check data
	//get response body
	bodyresp, err := ioutil.ReadAll(w.Body)
	rtstr := string(bodyresp)
	//decode data

	rtstr = mycrypto.DecodeOld(rtstr, 8)
	json.Unmarshal([]byte(rtstr), &rs)
	return
}

func setup() {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and
	// register your routes
	r = gin.Default()
	r.POST("/*name", postHandler)
}
func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}

//test special char
func TestSpecialChar(t *testing.T) {
	fmt.Println("==== test TestSpecialChar ====")
	rs, errstr := decodeResponse(mycrypto.EncDat2(c3mcommon.GetSpecialChar()), "")
	//check test data
	if rs.Status == "1" || errstr != "" {
		t.Fatalf("Test fail: %s - %s", rs.Message, errstr)
	}
	fmt.Printf("Pass: %+v\n", rs)
}

//test function
func TestCreateSex(t *testing.T) {
	fmt.Println("==== test TestCreateSex ====")
	rs, errstr := decodeResponse(mycrypto.EncDat2("CreateSex"), "")
	//check test data
	if rs.Status != "1" || errstr != "" {
		t.Fatalf("Service Response error: %s - %s", rs.Error, errstr)
	}
	fmt.Printf("Pass: %+v\n", rs)
}

//double test create session
func TestCreateSex2(t *testing.T) {
	fmt.Println("==== test TestCreateSex2 ====")
	rs, errstr := decodeResponse(mycrypto.EncDat2("CreateSex"), "")
	//check test data
	if rs.Status != "1" || errstr != "" {
		t.Fatalf("Service Response error: %s - %s", rs.Error, errstr)
	}
	testsession = rs.Message
	fmt.Printf("Pass: %+v\n", rs)
}

//test login
func TestLoginWithouSession(t *testing.T) {
	fmt.Println("==== test TestLoginWithouSession ====")
	data := "l|admin,123456"

	rs, errstr := decodeResponse(mycrypto.EncDat2("aut"), data)
	//check test data
	if rs.Status != "-2" || errstr != "" {
		t.Fatalf("Login test fail: %s - %s", rs.Error, errstr)
	}
	testsession = rs.Message
	fmt.Printf("Pass: %+v\n", rs)
}
func TestLoginWrongUser(t *testing.T) {
	fmt.Println("==== test TestLoginWrongUser ====")
	data := "l|admin,123456"

	rs, errstr := decodeResponse("aut|"+testsession, data)
	//check test data
	if rs.Status != "-2" || errstr != "" {
		t.Fatalf("Login test fail: %s - %s", rs.Error, errstr)
	}
	testsession = rs.Message
	fmt.Printf("Pass: %+v\n", rs)
}
func TestLoginSuccessUser(t *testing.T) {

}
