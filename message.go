package workwx

import (
	"fmt"
)

// Message 消息接口
type Message interface {
	SendAble() bool
}

// SendMessageResp 发送应用消息返回
type SendMessageResp struct {
	CommonResp
	InvalidUser    string `json:"invaliduser"`
	InvalidParty   string `json:"invalidparty"`
	UnlicensedUser string `json:"unlicenseduser"`
	MsgId          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}

// CommonMessage 通用消息结构
type CommonMessage struct {
	ToUser                 string `json:"touser,omitempty"`
	ToParty                string `json:"toparty,omitempty"`
	ToTag                  string `json:"totag,omitempty"`
	MsgType                string `json:"msgtype"`
	AgentId                int64  `json:"agentid"`
	EnableDuplicateCheck   uint8  `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval uint8  `json:"duplicate_check_interval,omitempty"`
}

// SafeMessage 保密消息
type SafeMessage struct {
	Safe uint8 `json:"safe"`
}

// IdTrans id转译
type IdTrans struct {
	EnableIdTrans int `json:"enable_id_trans,omitempty"`
}

type CommonText struct {
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// TextMessage 文本消息
type TextMessage struct {
	CommonMessage
	SafeMessage
	IdTrans
	CommonText
}

type CommonImage struct {
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

// ImageMessage 图片消息
type ImageMessage struct {
	CommonMessage
	SafeMessage
	CommonImage
}

type CommonVoice struct {
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

// VoiceMessage 语音消息
type VoiceMessage struct {
	CommonMessage
	CommonVoice
}

type CommonVideo struct {
	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
}

// VideoMessage 视频消息
type VideoMessage struct {
	CommonMessage
	SafeMessage
	CommonVideo
}

type CommonFile struct {
	File struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

// FileMessage 文件消息
type FileMessage struct {
	CommonMessage
	SafeMessage
	CommonFile
}

type CommonTextCard struct {
	TextCard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		BtnTxt      string `json:"btntxt"`
	} `json:"textcard"`
}

// TextCardMessage 文本卡片消息
type TextCardMessage struct {
	CommonMessage
	IdTrans
	CommonTextCard
}

type CommonNews struct {
	News struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			PicUrl      string `json:"picurl"`
			AppId       string `json:"appid"`
			PagePath    string `json:"pagepath"`
		} `json:"articles"`
	} `json:"news"`
}

// NewsMessage 图文消息
type NewsMessage struct {
	CommonMessage
	IdTrans
	CommonNews
}

type CommonMpNews struct {
	MpNews struct {
		Articles []struct {
			Title            string `json:"title"`
			ThumbMediaId     string `json:"thumb_media_id"`
			Author           string `json:"author"`
			ContentSourceUrl string `json:"content_source_url"`
			Content          string `json:"content"`
			Digest           string `json:"digest"`
		} `json:"articles"`
	} `json:"mpnews"`
}

// MpNewsMessage 图文消息
type MpNewsMessage struct {
	CommonMessage
	SafeMessage
	IdTrans
	CommonMpNews
}

type CommonMarkdown struct {
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

// MarkdownMessage markdown消息
type MarkdownMessage struct {
	CommonMessage
	CommonMarkdown
}

// SendMessage 发送应用消息
func (c *App) SendMessage(m Message) (*SendMessageResp, *Error) {
	if !m.SendAble() {
		return nil, NewError(10001, "invalid message type")
	}
	token, err := c.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("/cgi-bin/message/send?access_token=%s", token)
	var result SendMessageResp
	err = c.httpPost(uri, m, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m CommonMessage) SendAble() bool {
	return true
}

func NewTextMessage(agentId int64, content string, safe uint8) TextMessage {
	m := TextMessage{}
	m.AgentId = agentId
	m.MsgType = "text"
	m.Text = struct {
		Content string `json:"content"`
	}{
		Content: content,
	}
	m.Safe = safe
	return m
}

func NewImageMessage(agentId int64, mediaId string, safe uint8) ImageMessage {
	m := ImageMessage{}
	m.AgentId = agentId
	m.MsgType = "image"
	m.Image = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}
	m.Safe = safe
	return m
}

func NewVoiceMessage(agentId int64, mediaId string) VoiceMessage {
	m := VoiceMessage{}
	m.AgentId = agentId
	m.MsgType = "voice"
	m.Voice = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}
	return m
}

func NewVideoMessage(agentId int64, mediaId, description string, title string, safe uint8) VideoMessage {
	m := VideoMessage{}
	m.AgentId = agentId
	m.MsgType = "video"
	m.Video = struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		MediaId:     mediaId,
		Title:       title,
		Description: description,
	}
	m.Safe = safe
	return m
}

func NewFileMessage(agentId int64, mediaId string, safe uint8) FileMessage {
	m := FileMessage{}
	m.AgentId = agentId
	m.MsgType = "file"
	m.File = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}
	m.Safe = safe
	return m
}

func NewMarkdownMessage(agentId int64, content string) MarkdownMessage {
	m := MarkdownMessage{}
	m.AgentId = agentId
	m.MsgType = "markdown"
	m.Markdown = struct {
		Content string `json:"content"`
	}{
		Content: content,
	}
	return m
}
