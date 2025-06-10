package workwx

import "fmt"

// Tag 标签
type Tag struct {
	TagName string `json:"tagname"`
	TagId   int64  `json:"tagid"`
}

// CreateTagResp 创建标签请求参数
type CreateTagResp struct {
	CommonResp
	TagId int64 `json:"tagid"`
}

// GetTagResp 获取标签及成员返回
type GetTagResp struct {
	CommonResp
	TagName  string `json:"tagname"`
	UserList []struct {
		UserId string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
	PartyList []int64 `json:"partylist"`
}

// TagUsers 标签成员请求参数
type TagUsers struct {
	TagId     int64    `json:"tagid"`
	UserList  []string `json:"userlist"`
	PartyList []int64  `json:"partylist"`
}

// TagUsersResp 增加/删除标签成员返回
type TagUsersResp struct {
	CommonResp
	InvalidList  string  `json:"invalidlist"`
	InvalidParty []int64 `json:"invalidparty"`
}

// GetTagListResp 获取标签列表返回
type GetTagListResp struct {
	CommonResp
	TagList []*Tag `json:"taglist"`
}

// CreateTag 创建标签
func (c *App) CreateTag(req *Tag) (int64, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return 0, err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/create?access_token=%s", token)
	var result CreateTagResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return 0, err
	}
	return result.TagId, nil
}

// UpdateTag 更新标签名字
func (c *App) UpdateTag(req *Tag) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/update?access_token=%s", token)
	var result CommonResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag 删除标签
func (c *App) DeleteTag(tagId int64) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/delete?access_token=%s&tagid=%d", token, tagId)
	var result CommonResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return err
	}
	return nil
}

// GetTag 获取标签成员
func (c *App) GetTag(tagId int64) (*GetTagResp, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/get?access_token=%s&tagid=%d", token, tagId)
	var result GetTagResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

// AddTagUsers 增加标签成员
func (c *App) AddTagUsers(req *TagUsers) (*TagUsersResp, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/addtagusers?access_token=%s", token)
	var result TagUsersResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DelTagUsers 删除标签成员
func (c *App) DelTagUsers(req *TagUsers) (*TagUsersResp, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/deltagusers?access_token=%s", token)
	var result TagUsersResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTagList 获取标签列表
func (c *App) GetTagList() ([]*Tag, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/tag/list?access_token=%s", token)
	var result GetTagListResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}
	return result.TagList, err
}
