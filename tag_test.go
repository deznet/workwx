package workwx

import (
	"os"
	"strconv"
	"testing"
)

func TestApp_CreateTag(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	req := new(Tag)
	req.TagId = 10086
	req.TagName = "测试标签"
	tagId, err := app.CreateTag(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tagId)
}

func TestApp_UpdateTag(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	req := new(Tag)
	req.TagId = 10086
	req.TagName = "测试标签2"
	err := app.UpdateTag(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_DeleteTag(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	err := app.DeleteTag(10086)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_AddTagUsers(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	tagUsers := new(TagUsers)
	tagUsers.TagId = 10086
	tagUsers.UserList = []string{os.Getenv("TEST_USERID")}
	resp, err := app.AddTagUsers(tagUsers)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_DelTagUsers(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	tagUsers := new(TagUsers)
	tagUsers.TagId = 10086
	tagUsers.UserList = []string{os.Getenv("TEST_USERID")}
	resp, err := app.DelTagUsers(tagUsers)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_GetTag(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	resp, err := app.GetTag(10086)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_GetTagList(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	resp, err := app.GetTagList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
