package workwx

import (
	"os"
	"strconv"
	"testing"
)

func TestApp_CreateAppChat(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	req := new(CreateAppChatReq)
	req.Name = "测试群组"
	req.UserList = []string{os.Getenv("TEST_USERID"), os.Getenv("TEST_USERID_2")}
	req.Owner = os.Getenv("TEST_USERID")
	req.ChatId = "abcdef123456"
	id, err := app.CreateAppChat(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func TestApp_UpdateAppChat(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	req := new(UpdateAppChatReq)
	req.Name = "测试群组1"
	req.AddUserList = []string{}
	req.DelUserList = []string{"kd_494806"}
	req.Owner = os.Getenv("TEST_USERID")
	req.ChatId = "abcdef123456"
	err := app.UpdateAppChat(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_GetAppChat(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	info, err := app.GetAppChat("abcdef123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestApp_SendAppChatTextMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatTextMessage{}
	m.ChatId = "abcdef123456"
	m.MsgType = "text"
	m.Text.Content = "这是一条测试消息"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_SendAppChatImageMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatImageMessage{}
	m.Image.MediaId = "2Rwj-Nb0IpxSxkJxmgKkPWdc1t8H-mmf21V92w3FXMWYA54YdvMlcbV_D4wwaVxLg"
	m.ChatId = "abcdef123456"
	m.MsgType = "image"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}

}

// 说明：发送到微信端语音有问题，企业微信端没问题
func TestApp_SendAppChatVoiceMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatVoiceMessage{}
	m.ChatId = "abcdef123456"
	m.MsgType = "voice"
	m.Voice.MediaId = "2cnEmOZzgSRcuadWzBrpm2bZqrqpWXa84yg7X42lasCiMVZf4H01XBi98-67KuCUV"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_SendAppChatVideoMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatVideoMessage{}
	m.ChatId = "abcdef123456"
	m.MsgType = "video"
	m.Video.MediaId = "2cnEmOZzgSRcuadWzBrpm2bZqrqpWXa84yg7X42lasCiMVZf4H01XBi98-67KuCUV"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_SendAppChatFileMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatFileMessage{}
	m.ChatId = "abcdef123456"
	m.MsgType = "file"
	m.File.MediaId = "2aFk9yf0h4OfbDJt4q24LweZDxSfHXMGffGHetbj4hy3BbbURNTG59XwjFe7mHnJ9"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_SendAppChatTextCardMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatTextCardMessage{}
	m.TextCard = struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		BtnTxt      string `json:"btntxt"`
	}{
		Title:       "测试卡片消息",
		Description: "测试卡片消息描述",
		URL:         "https://www.qq.com",
		BtnTxt:      "点击进入",
	}
	m.ChatId = "abcdef123456"
	m.MsgType = "textcard"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}

}

func TestApp_SendAppChatNewsMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatNewsMessage{}
	m.News = struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			PicUrl      string `json:"picurl"`
			AppId       string `json:"appid"`
			PagePath    string `json:"pagepath"`
		} `json:"articles"`
	}{
		Articles: []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			PicUrl      string `json:"picurl"`
			AppId       string `json:"appid"`
			PagePath    string `json:"pagepath"`
		}{
			{
				Title:       "人工智能技术新突破",
				Description: "科学家成功开发出可自主学习的新型AI模型，准确率达98%",
				URL:         "https://qq.com",
				PicUrl:      "https://inews.gtimg.com/om_bt/OiVKxHO4yTCg_711Y0Mk8RO6OMFZx410mnS2MUX1fXRxcAA/641",
				AppId:       "",
				PagePath:    "",
			},
			{
				Title:       "全球环保倡议取得进展",
				Description: "联合国报告显示：2025年碳排放量首次实现负增长",
				URL:         "https://qq.com",
				PicUrl:      "https://inews.gtimg.com/om_bt/OiVKxHO4yTCg_711Y0Mk8RO6OMFZx410mnS2MUX1fXRxcAA/641",
				AppId:       "",
				PagePath:    "",
			},
		},
	}

	m.MsgType = "news"
	m.ChatId = "abcdef123456"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_SendAppChatMpNewsMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := AppChatMpNewsMessage{}
	m.MpNews = struct {
		Articles []struct {
			Title            string `json:"title"`
			ThumbMediaId     string `json:"thumb_media_id"`
			Author           string `json:"author"`
			ContentSourceUrl string `json:"content_source_url"`
			Content          string `json:"content"`
			Digest           string `json:"digest"`
		} `json:"articles"`
	}{
		Articles: []struct {
			Title            string `json:"title"`
			ThumbMediaId     string `json:"thumb_media_id"`
			Author           string `json:"author"`
			ContentSourceUrl string `json:"content_source_url"`
			Content          string `json:"content"`
			Digest           string `json:"digest"`
		}{
			{
				Title:            "人工智能技术新突破",
				ThumbMediaId:     "2Rwj-Nb0IpxSxkJxmgKkPWdc1t8H-mmf21V92w3FXMWYA54YdvMlcbV_D4wwaVxLg",
				Author:           "qq",
				ContentSourceUrl: "https://qq.com",
				Content:          "6月9日，苏州市人工智能行业协会称，计划征集人工智能技术赋能苏州足球队的创新产品及解决方案，促进提升训练水平和竞技表现，助力苏州足球队在2025年江苏省城市足球联赛中争创佳绩。",
				Digest:           "6月9日，苏州市人工智能行业协会称，计划征集人工智能技术赋能苏州足球队的创新产品及解决方案",
			},
		},
	}
	m.MsgType = "mpnews"
	m.ChatId = "abcdef123456"
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApp_SendAppChatMarkdownMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	content := "您的会议室已经预定，稍后会同步到`邮箱`  \n>**事项详情**  \n>事　项：<font color=\"info\">开会</font>  \n>组织者：@miglioguan  \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang  \n>  \n>会议室：<font color=\"info\">广州TIT 1楼 301</font>  \n>日　期：<font color=\"warning\">2018年5月18日</font>  \n>时　间：<font color=\"comment\">上午9:00-11:00</font>  \n>  \n>请准时参加会议。  \n>  \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"
	m := AppChatMarkdownMessage{}
	m.MsgType = "markdown"
	m.ChatId = "abcdef123456"
	m.Markdown.Content = content
	err := app.SendAppChatMessage(m)
	if err != nil {
		t.Fatal(err)
	}
}
