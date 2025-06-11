package workwx

import (
	"github.com/go-resty/resty/v2"
)

// ChatBot 群机器人
type ChatBot struct {
	WebHookUrl string
}

type ChatBotMessage interface {
	ChatBotSendAble() bool
}

type ChatBotCommonMessage struct {
	MsgType string `json:"msgtype"`
}

type ChatBotTextMessage struct {
	ChatBotCommonMessage
	Text struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

type ChatBotMarkdownMessage struct {
	ChatBotCommonMessage
	CommonMarkdown
}

type ChatBotImageMessage struct {
	ChatBotCommonMessage
	Image struct {
		Base64 string `json:"base64"`
		Md5    string `json:"md5"`
	} `json:"image"`
}

type ChatBotNewsMessage struct {
	ChatBotCommonMessage
	CommonNews
}

type ChatBotFileMessage struct {
	ChatBotCommonMessage
	CommonFile
}

type ChatBotVoiceMessage struct {
	ChatBotCommonMessage
	CommonVoice
}

func NewChatBot(webHookUrl string) *ChatBot {
	return &ChatBot{WebHookUrl: webHookUrl}
}

func (b *ChatBot) Send(m ChatBotMessage) *Error {
	if !m.ChatBotSendAble() {
		return NewError(10001, "invalid message type")
	}
	httpClient := resty.New()
	request := httpClient.R()
	request.SetHeader("Accept", "application/json")
	var result CommonResp
	request.SetResult(&result)
	request.SetBody(m)
	_, err := request.Post(b.WebHookUrl)
	if err != nil {
		return NewError(10000, err.Error())
	}
	if result.IsOK() {
		return nil
	}
	return NewError(result.GetErrorCode(), result.GetError().Error())
}

func (c ChatBotCommonMessage) ChatBotSendAble() bool {
	return true
}
