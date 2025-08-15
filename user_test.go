package workwx

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
)

func TestApp_GetUser(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	user, err := app.GetUser(os.Getenv("TEST_USERID"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestApp_CreateUser(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	user := User{}
	user.UserID = os.Getenv("TEST_USERID")
	user.Name = os.Getenv("TEST_USERNAME")
	user.Department = []int{11090}
	user.Mobile = os.Getenv("TEST_USER_MOBILE")
	user.Enable = 1
	user.Gender = "1"
	err := app.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_UpdateUser(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	user := User{}
	user.UserID = os.Getenv("TEST_USERID")
	user.Name = os.Getenv("TEST_USERNAME")
	user.Department = []int{11090}
	user.Mobile = os.Getenv("TEST_USER_MOBILE")
	user.Enable = 1
	user.Gender = "1"
	user.Email = os.Getenv("TEST_USER_EMAIL")
	err := app.UpdateUser(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_DeleteUser(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	err := app.DeleteUser(os.Getenv("TEST_USERID"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_BatchDeleteUser(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	err := app.BatchDeleteUser([]string{os.Getenv("TEST_USERID")})
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_UserIdToOpenId(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	openId, err := app.UserIdToOpenId(os.Getenv("TEST_USERID"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(openId)
}

func TestApp_OpenIdToUserId(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	openId, err := app.OpenIdToUserId(os.Getenv("TEST_OPENID"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(openId)
}

func TestApp_GetJoinQrcode(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	qr, err := app.GetJoinQrcode(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(qr)
}

func TestApp_GetUserIdByMobile(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	userid, err := app.GetUserIdByMobile(os.Getenv("TEST_USER_MOBILE"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userid)
}

func TestApp_GetUserIdByEmail(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	req := new(GetUserIdByEmailReq)
	req.EmailType = 2
	req.Email = os.Getenv("TEST_USER_EMAIL")
	userid, err := app.GetUserIdByEmail(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userid)
}

func TestApp_GetUserIdList(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	req := new(GetUserIdListReq)
	req.Cursor = ""
	req.Limit = 10
	resp, err := app.GetUserIdList(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
