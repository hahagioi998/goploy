package test

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/zhenorzz/goploy/config"
//	"github.com/zhenorzz/goploy/core"
//	"github.com/zhenorzz/goploy/model"
//	"github.com/zhenorzz/goploy/route"
//	"math/rand"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"time"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func TestApi(t *testing.T) {
//	config.Create(core.GetConfigFile())
//	core.CreateValidator()
//	model.Init()
//	// <setup code>
//	userLogin(t)
//	addUser(t)
//	addServer(t)
//	// <logic code>
//	// user
//	t.Run("user/info", userInfo)
//	t.Run("user/getList", getUserList)
//	t.Run("user/getOption", getUserOption)
//	t.Run("user/edit", editUser)
//	t.Run("user/changePassword", changeUserPassword)
//
//	// server
//	t.Run("server/getList", getServerList)
//	t.Run("server/getInstallPreview", GetServerInstallPreview)
//	t.Run("server/getInstallList", getServerInstallList)
//	t.Run("server/getOption", getServerOption)
//	t.Run("server/edit", editServer)
//
//	// <tear-down code>
//	removeUser(t)
//	removeServer(t)
//}
//
//var handler = route.Init()
//var (
//	// set when user login
//	token string
//	// set when user is added
//	userID int64
//	// set when server is added
//	serverID int64
//)
//
//func request(t *testing.T, method, url string, body interface{}) core.Response {
//	buf := new(bytes.Buffer)
//	_ = json.NewEncoder(buf).Encode(body)
//	req, err := http.NewRequest(method, url, buf)
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
//	if token != "" {
//		req.Header.Set("Cookie", fmt.Sprintf("%s=%s", config.Toml.Cookie.Name, token))
//	}
//	r := httptest.NewRecorder()
//	handler.ServeHTTP(r, req)
//	// 检测返回的状态码
//	if r.Code != http.StatusOK {
//		t.Fatalf("http request error, code: %d", r.Code)
//	}
//
//	var resp core.Response
//
//	// 检测返回的json格式
//	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
//		t.Fatal(err.Error())
//	}
//
//	// 检测接口返回值
//	if resp.Code != core.Pass {
//		t.Fatalf("http response error, content: %v", resp)
//	}
//	return resp
//}
//
//func getRandomStringOf(length int) string {
//	str := "0123456789abcdefghijklmnopqrstuvwxyz"
//	bs := []byte(str)
//	var result []byte
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for i := 0; i < length; i++ {
//		result = append(result, bs[r.Intn(len(bs))])
//	}
//	return string(result)
//}
