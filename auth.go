package workwx

import (
	"fmt"
	"github.com/jinzhu/copier"
	"net/url"
	"strconv"
)

// GetOAuth2UrlReq 构造网页授权链接请求链接
type GetOAuth2UrlReq struct {
	RedirectUri string
	Scope       string
	State       string
	AgentId     int64
}

// AuthUserInfo 鉴权用户信息
type AuthUserInfo struct {
	//成员UserID,用户为企业成员时返回
	UserID string `json:"userid"`

	//成员票据，最大为512字节，有效期为1800s,用户为企业成员时返回
	UserTicket string `json:"user_ticket"`

	//非企业成员的标识，对当前企业唯一。不超过64字节。用户为非企业成员时返回
	OpenID string `json:"Openid"`

	//外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回。用户为非企业成员时返回
	ExternalUserID string `json:"external_userid"`
}

// GetUserInfoByAuthCodeResp 通过code获取用户信息返回信息
type GetUserInfoByAuthCodeResp struct {
	CommonResp
	AuthUserInfo
}

// GetAuthUserDetailReq 获取访问用户敏感信息请求包体
type GetAuthUserDetailReq struct {
	UserTicket string `json:"user_ticket"`
}

// AuthUserDetail 用户敏感信息
type AuthUserDetail struct {
	UserId   string `json:"userid"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	QrCode   string `json:"qr_code"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	BizEmail string `json:"biz_mail"`
	Address  string `json:"address"`
}

// GetAuthUserDetailResp 获取访问用户敏感信息返回结果
type GetAuthUserDetailResp struct {
	CommonResp
	AuthUserDetail
}

// GetOAuth2Url 构造网页授权链接
func (c *App) GetOAuth2Url(req *GetOAuth2UrlReq) string {
	u, _ := url.Parse("https://open.weixin.qq.com/connect/oauth2/authorize")
	q := u.Query()
	q.Add("appid", c.CorpID)
	q.Add("redirect_uri", req.RedirectUri)
	q.Add("response_type", "code")
	q.Add("scope", req.Scope)
	q.Add("state", req.State)
	q.Add("agentid", strconv.FormatInt(req.AgentId, 10))
	u.RawQuery = q.Encode()
	u.Fragment = "wechat_redirect"
	return u.String()
}

// GetUserInfoByAuthCode 通过code获取用户信息
func (c *App) GetUserInfoByAuthCode(code string) (*AuthUserInfo, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/auth/getuserinfo?access_token=%s&code=%s", token, code)
	var result GetUserInfoByAuthCodeResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}

	userInfo := new(AuthUserInfo)
	err1 := copier.Copy(userInfo, &result)
	if err1 != nil {
		return nil, NewError(10000, err.Error())
	}
	return userInfo, nil
}

// GetAuthUserDetail 获取访问用户敏感信息
func (c *App) GetAuthUserDetail(userTicket string) (*AuthUserDetail, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/auth/getuserdetail?access_token=%s", token)
	var result GetAuthUserDetailResp
	req := GetAuthUserDetailReq{
		UserTicket: userTicket,
	}
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return nil, err
	}
	userDetail := new(AuthUserDetail)
	err1 := copier.Copy(userDetail, &result)
	if err1 != nil {
		return nil, NewError(10000, err1.Error())
	}
	return userDetail, nil
}
