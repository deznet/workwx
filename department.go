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

// GetDepartmentListResp 获取子部门ID列表返回
type GetDepartmentListResp struct {
	CommonResp
	Departments []*Department `json:"department_id"`
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

// GetDepartmentList 获取子部门ID列表,只能返回id\parent_id\order
func (c *App) GetDepartmentList(id uint32) ([]*Department, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/department/simplelist?access_token=%s&id=%d", token, id)
	var result GetDepartmentListResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}
	return result.Departments, nil
}

// GetDepartment 获取部门信息，只能返回id\parent_id\order
func (c *App) GetDepartment(id uint32) (*Department, *Error) {
	list, err := c.GetDepartmentList(id)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		if v.Id == id {
			return v, nil
		}
	}
	return &Department{}, nil
}
