package workwx

import "fmt"

// Department 部门
type Department struct {
	Name     string `json:"name"`
	NameEn   string `json:"name_en,omitempty"`
	ParentId uint32 `json:"parentid"`
	Order    uint32 `json:"order"`
	Id       uint32 `json:"id,omitempty"`
}

// CreateDepartmentResp 创建部门返回
type CreateDepartmentResp struct {
	CommonResp
	Id uint32 `json:"id"`
}

// GetDepartmentResp 获取单个部门详情返回
type GetDepartmentResp struct {
	CommonResp
	Department *Department `json:"department"`
}

// CreateDepartment 创建部门
func (c *App) CreateDepartment(req *Department) (uint32, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return 0, err
	}
	uri := fmt.Sprintf("/cgi-bin/department/create?access_token=%s", token)
	var result CreateDepartmentResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}

// UpdateDepartment 更新部门
func (c *App) UpdateDepartment(req *Department) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/department/update?access_token=%s", token)
	var result CommonResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDepartment 删除部门
func (c *App) DeleteDepartment(id uint32) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/department/delete?access_token=%s&id=%d", token, id)
	var result CommonResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return err
	}
	return nil
}

// GetDepartment 获取单个部门详情
// 已不可调用
func (c *App) GetDepartment(id uint32) (*Department, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/department/get?access_token=%s&id=%d", token, id)
	var result GetDepartmentResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}
	return result.Department, nil
}
