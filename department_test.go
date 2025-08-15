package workwx

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
)

func TestApp_CreateDepartment(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	req := new(Department)
	req.Name = "测试部门"
	req.ParentId = 1
	id, err := app.CreateDepartment(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestApp_UpdateDepartment(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	req := new(Department)
	req.Name = "测试部门1"
	req.ParentId = 1
	req.Id = 4000001
	err := app.UpdateDepartment(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_DeleteDepartment(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	err := app.DeleteDepartment(4000001)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_GetDepartmentList(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	depts, err := app.GetDepartmentList(4000000)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, v := range depts {
		t.Log(v)
	}
}

func TestApp_GetDepartment(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("CONTACT_APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("CONTACT_APP_SECRET"))
	dept, err := app.GetDepartment(4000000)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(dept)
}
