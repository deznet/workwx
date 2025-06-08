package workwx

import "fmt"

// User 成员信息
type User struct {
	UserID              string   `json:"userid"`
	Name                string   `json:"name"`
	Alias               string   `json:"alias"`
	Mobile              string   `json:"mobile"`
	Department          []int    `json:"department"`
	Order               []int    `json:"order"`
	Position            string   `json:"position"`
	Gender              uint8    `json:"gender"`
	Email               string   `json:"email"`
	BizMail             string   `json:"biz_mail"`
	IsLeaderInDept      []int    `json:"is_leader_in_dept"`
	DirectLeader        []string `json:"direct_leader"`
	Telephone           string   `json:"telephone"`
	Avatar              string   `json:"avatar"`       //读取成员才返回
	ThumbAvatar         string   `json:"thumb_avatar"` //读取成员才返回
	Address             string   `json:"address"`
	OpenUserId          string   `json:"open_userid"` //读取成员才返回
	MainDepartment      int      `json:"main_department"`
	Status              int      `json:"status"`  //读取成员才返回
	QrCode              string   `json:"qr_code"` //读取成员才返回
	ExternalPosition    string   `json:"external_position"`
	UserExtAttr         `json:"extattr"`
	UserExternalProfile `json:"external_profile"`
	//以下添加更新操作才有此参数
	Enable        int    `json:"enable"`
	AvatarMediaId string `json:"avatarMediaId"`
	//添加成员才有此参数
	ToInvite bool `json:"to_invite"`
}

// UserExtAttr 用户扩展属性
type UserExtAttr struct {
	Attrs []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Text struct {
			Value string `json:"value"`
		} `json:"text,omitempty"`
		Web struct {
			URL   string `json:"url"`
			Title string `json:"title"`
		} `json:"web,omitempty"`
	} `json:"attrs"`
}

// UserExternalProfile 成员对外属性
type UserExternalProfile struct {
	ExternalCorpName string `json:"external_corp_name"`
	WechatChannels   struct {
		NickName string `json:"nickname"`
		Status   int    `json:"status"`
	} `json:"wechat_channels"`
	ExternalAttr []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Text struct {
			Value string `json:"value"`
		} `json:"text,omitempty"`
		Web struct {
			URL   string `json:"url"`
			Title string `json:"title"`
		} `json:"web,omitempty"`
		MiniProgram struct {
			Appid    string `json:"appid"`
			PagePath string `json:"pagepath"`
			Title    string `json:"title"`
		} `json:"miniprogram,omitempty"`
	} `json:"external_attr"`
}

// GetUserResp 获取成员信息
type GetUserResp struct {
	CommonResp
	User
}

// BatchDeleteUserReq 批量删除成员
type BatchDeleteUserReq struct {
	UserIDList []string `json:"useridlist"`
}

// UserIdToOpenIdReq userid转openid请求参数
type UserIdToOpenIdReq struct {
	UserId string `json:"userid"`
}

// UserIdToOpenIdResp userid转openid返回结果
type UserIdToOpenIdResp struct {
	CommonResp
	OpenId string `json:"openid"`
}

// OpenIdToUserIdReq openid转userid请求参数
type OpenIdToUserIdReq struct {
	OpenId string `json:"openid"`
}

// OpenIdToUserIdResp openid转userid返回结果
type OpenIdToUserIdResp struct {
	CommonResp
	UserId string `json:"userid"`
}

// GetJoinQrcodeResp 获取加入企业二维码返回结果
type GetJoinQrcodeResp struct {
	CommonResp
	JoinQrcode string `json:"join_qrcode"`
}

// GetUserIdByMobileReq 手机号获取userid请求参数
type GetUserIdByMobileReq struct {
	Mobile string `json:"mobile"`
}

// GetUserIdByMobileResp 手机号获取userid返回
type GetUserIdByMobileResp struct {
	CommonResp
	UserId string `json:"userid"`
}

// GetUserIdByEmailReq 邮箱获取userid请求参数
type GetUserIdByEmailReq struct {
	Email     string `json:"email"`
	EmailType uint8  `json:"email_type"`
}

// GetUserIdByEmailResp 邮箱获取userid返回
type GetUserIdByEmailResp struct {
	CommonResp
	UserId string `json:"userid"`
}

// GetUserIdListReq 获取成员ID列表请求参数
type GetUserIdListReq struct {
	Cursor string `json:"cursor"`
	Limit  uint32 `json:"limit"`
}

// GetUserIdListResp 获取成员ID列表返回
type GetUserIdListResp struct {
	CommonResp
	NextCursor string      `json:"next_cursor"`
	DeptUsers  []*DeptUser `json:"dept_user"`
}

// DeptUser 用户-部门关系列表
type DeptUser struct {
	UserId     string `json:"userid"`
	Department uint32 `json:"department"`
}

// GetUser 读取成员
func (c *App) GetUser(userID string) (*GetUserResp, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/user/get?access_token=%s&userid=%s", token, userID)
	result := new(GetUserResp)
	err = c.httpGet(uri, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser 创建成员
func (c *App) CreateUser(req User) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/user/create?access_token=%s", token)
	var result CommonResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户
func (c *App) UpdateUser(req User) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/user/update?access_token=%s", token)
	var result CommonResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func (c *App) DeleteUser(userID string) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/user/delete?access_token=%s&userid=%s", token, userID)
	var result CommonResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return err
	}
	return nil
}

// BatchDeleteUser 批量删除成员
func (c *App) BatchDeleteUser(userIdList []string) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/user/batchdelete?access_token=%s", token)
	var result CommonResp
	req := new(BatchDeleteUserReq)
	req.UserIDList = userIdList
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return err
	}
	return nil
}

// UserIdToOpenId userid转openid
func (c *App) UserIdToOpenId(userId string) (string, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("/cgi-bin/user/convert_to_openid?access_token=%s", token)
	var result UserIdToOpenIdResp
	req := new(UserIdToOpenIdReq)
	req.UserId = userId
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return "", err
	}
	return result.OpenId, nil
}

// OpenIdToUserId openid转userid
func (c *App) OpenIdToUserId(openId string) (string, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("/cgi-bin/user/convert_to_userid?access_token=%s", token)
	var result OpenIdToUserIdResp
	req := new(OpenIdToUserIdReq)
	req.OpenId = openId
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return "", err
	}
	return result.UserId, nil
}

// GetJoinQrcode 获取加入企业二维码
func (c *App) GetJoinQrcode(sizeType uint8) (string, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("/cgi-bin/corp/get_join_qrcode?access_token=%s&size_type=%d", token, sizeType)
	var result GetJoinQrcodeResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return "", err
	}
	return result.JoinQrcode, nil
}

// GetUserIdByMobile 手机号获取userid
func (c *App) GetUserIdByMobile(mobile string) (string, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("/cgi-bin/user/getuserid?access_token=%s", token)
	var result GetUserIdByMobileResp
	var req GetUserIdByMobileReq
	req.Mobile = mobile
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return "", err
	}
	return result.UserId, nil
}

// GetUserIdByEmail 邮箱获取userid
func (c *App) GetUserIdByEmail(req *GetUserIdByEmailReq) (string, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("/cgi-bin/user/get_userid_by_email?access_token=%s", token)
	var result GetUserIdByEmailResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return "", err
	}
	return result.UserId, nil
}

// GetUserIdList 获取成员ID列表
func (c *App) GetUserIdList(req *GetUserIdListReq) (*GetUserIdListResp, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/user/list_id?access_token=%s", token)
	var result GetUserIdListResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
