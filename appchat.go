package workwx

import "fmt"

// CreateAppChatReq 创建群聊会话请求参数
type CreateAppChatReq struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
	ChatId   string   `json:"chatid"`
}

// CreateAppChatResp 创建群聊会话返回
type CreateAppChatResp struct {
	CommonResp
	ChatId string `json:"chatid"`
}

// UpdateAppChatReq 修改群聊会话请求参数
type UpdateAppChatReq struct {
	ChatId      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list,omitempty"`
	DelUserList []string `json:"del_user_list,omitempty"`
}

// ChatInfo 群聊会话
type ChatInfo struct {
	ChatId   string   `json:"chatid"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	UserList []string `json:"userlist"`
}

// GetAppChatResp  获取群聊会话返回
type GetAppChatResp struct {
	CommonResp
	ChatInfo *ChatInfo `json:"chat_info"`
}

type AppChatMessage interface {
	AppChatSendAble() bool
}

type AppChatCommonMessage struct {
	ChatId  string `json:"chatid"`
	MsgType string `json:"msgtype"`
}

// AppChatTextMessage 文本消息
type AppChatTextMessage struct {
	AppChatCommonMessage
	SafeMessage
	CommonText
}

// AppChatImageMessage 图片消息
type AppChatImageMessage struct {
	AppChatCommonMessage
	SafeMessage
	CommonImage
}

// AppChatVoiceMessage 语音消息
type AppChatVoiceMessage struct {
	AppChatCommonMessage
	CommonVoice
}

// AppChatVideoMessage 视频消息
type AppChatVideoMessage struct {
	AppChatCommonMessage
	SafeMessage
	CommonVideo
}

// AppChatFileMessage 文件消息
type AppChatFileMessage struct {
	AppChatCommonMessage
	SafeMessage
	CommonFile
}

// AppChatTextCardMessage 文本卡片消息
type AppChatTextCardMessage struct {
	AppChatCommonMessage
	CommonTextCard
}

// AppChatNewsMessage 图文消息
type AppChatNewsMessage struct {
	AppChatCommonMessage
	CommonNews
}

// AppChatMpNewsMessage 图文消息
type AppChatMpNewsMessage struct {
	AppChatCommonMessage
	CommonMpNews
}

// AppChatMarkdownMessage markdown消息
type AppChatMarkdownMessage struct {
	AppChatCommonMessage
	CommonMarkdown
}

// CreateAppChat 创建群聊会话
func (c *App) CreateAppChat(req *CreateAppChatReq) (string, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("/cgi-bin/appchat/create?access_token=%s", token)
	var result CreateAppChatResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return "", err
	}
	return result.ChatId, nil
}

// UpdateAppChat 修改群聊会话
func (c *App) UpdateAppChat(req *UpdateAppChatReq) *Error {
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/appchat/update?access_token=%s", token)
	var result CommonResp
	err = c.httpPost(uri, req, &result)
	if err != nil {
		return err
	}
	return nil
}

// GetAppChat 获取群聊会话
func (c *App) GetAppChat(chatId string) (*ChatInfo, *Error) {
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/appchat/get?access_token=%s&chatid=%s", token, chatId)
	var result GetAppChatResp
	err = c.httpGet(uri, &result)
	if err != nil {
		return nil, err
	}
	return result.ChatInfo, nil
}

// SendAppChatMessage 应用推送消息到群聊会话
func (c *App) SendAppChatMessage(m AppChatMessage) *Error {
	if !m.AppChatSendAble() {
		return NewError(10001, "invalid message type")
	}
	token, err := c.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/cgi-bin/appchat/send?access_token=%s", token)
	var result CommonResp
	err = c.httpPost(uri, m, &result)
	if err != nil {
		return err
	}
	return nil
}

func (m AppChatCommonMessage) AppChatSendAble() bool {
	return true
}
