package workwx

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
)

func TestApp_GetOAuth2Url(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	agentId, _ := strconv.ParseInt(os.Getenv("AUTH_AGENTID"), 10, 64)
	uri := app.GetOAuth2Url(&GetOAuth2UrlReq{
		RedirectUri: os.Getenv("REDIRECT_URI"),
		Scope:       os.Getenv("AUTH_SCOPE"),
		State:       "12345",
		AgentId:     agentId,
	})
	t.Log(uri)
}

func TestApp_GetUserInfoByAuthCode(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	userinfo, err := app.GetUserInfoByAuthCode(os.Getenv("AUTH_CODE"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(userinfo)
}

func TestApp_GetAuthUserDetail(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	data, err := app.GetAuthUserDetail(os.Getenv("AUTH_USER_TICKET"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
