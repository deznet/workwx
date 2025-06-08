package workwx

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"sync"
	"time"
)

// WorkWx 企业微信客户端
type WorkWx struct {
	// 企业 ID
	CorpID string
}

// App 应用，企业微信每个应用是分开的
type App struct {
	*WorkWx
	AppID           int64
	AppSecret       string
	accessToken     *AccessToken
	accessTokenLock *sync.RWMutex
}

type AccessToken struct {
	Token     string
	ExpiresAt time.Time
}

type ICommonResp interface {
	IsOK() bool
	GetError() error
	GetErrorCode() int
}

// CommonResp 企业微信返回参数公共部分
type CommonResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Error 企业微信错误，区别于系统error
type Error struct {
	Code int
	Msg  string
}

// GetTokenResp 企业微信获取token返回参数
type GetTokenResp struct {
	CommonResp
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// New 新建企业微信客户端
func New(corpID string) *WorkWx {
	return &WorkWx{
		CorpID: corpID,
	}
}

// WithApp 创建企业微信应用客户端
func (c *WorkWx) WithApp(appID int64, appSecret string) *App {
	return &App{
		WorkWx:          c,
		AppID:           appID,
		AppSecret:       appSecret,
		accessTokenLock: &sync.RWMutex{},
		accessToken:     &AccessToken{},
	}
}

// GetAccessToken 获取AccessToken
func (c *App) GetAccessToken() (string, error) {
	//根据accessToken的有效期判断
	if c.accessToken.ExpiresAt.After(time.Now()) {
		return c.accessToken.Token, nil
	}
	//从企业微信服务器获取
	token, err := c.GetAccessTokenFromServer()
	if err != nil {
		return "", err
	}
	c.accessToken = token
	return c.accessToken.Token, nil
}

func (c *App) GetAccessTokenFromServer() (*AccessToken, error) {
	uri := fmt.Sprintf("/cgi-bin/gettoken?corpid=%s&corpsecret=%s", c.CorpID, c.AppSecret)
	var result GetTokenResp
	err := c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}
	token := new(AccessToken)
	token.Token = result.Token
	now := time.Now()
	expiresIn := time.Duration(result.ExpiresIn - 1800)
	token.ExpiresAt = now.Add(expiresIn * time.Second)
	return token, nil
}

// httpGet 企业微信请求接口get通用方法
func (c *App) httpGet(uri string, resp ICommonResp) error {
	httpClient := resty.New()
	httpClient.BaseURL = "https://qyapi.weixin.qq.com"
	request := httpClient.R()
	request.SetHeader("Accept", "application/json")
	request.SetResult(&resp)
	_, err := request.Get(uri)
	if err != nil {
		return err
	}
	if resp.IsOK() {
		return nil
	}
	return NewError(resp.GetErrorCode(), resp.GetError().Error())
}

// httpPost 企业微信请求接口post通用方法
func (c *App) httpPost(uri string, body interface{}, resp ICommonResp) error {
	httpClient := resty.New()
	httpClient.BaseURL = "https://qyapi.weixin.qq.com"
	request := httpClient.R()
	request.SetHeader("Accept", "application/json")
	request.SetResult(&resp)
	request.SetBody(body)
	_, err := request.Post(uri)
	if err != nil {
		return err
	}
	if resp.IsOK() {
		return nil
	}
	return NewError(resp.GetErrorCode(), resp.GetError().Error())
}

func (x *CommonResp) IsOK() bool {
	return x.ErrCode == 0
}

func (x *CommonResp) GetError() error {
	return errors.New(x.ErrMsg)
}

func (x *CommonResp) GetErrorCode() int {
	return x.ErrCode
}
