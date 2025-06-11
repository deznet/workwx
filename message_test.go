package workwx

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"testing"
)

func TestApp_SendTextMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := NewTextMessage(appId, "这是一条测试消息", 0)
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendImageMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := NewImageMessage(appId, "2Rwj-Nb0IpxSxkJxmgKkPWdc1t8H-mmf21V92w3FXMWYA54YdvMlcbV_D4wwaVxLg", 0)
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

// 说明：发送到微信端语音有问题，企业微信端没问题
func TestApp_SendVoiceMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := NewVoiceMessage(appId, "2cnEmOZzgSRcuadWzBrpm2bZqrqpWXa84yg7X42lasCiMVZf4H01XBi98-67KuCUV")
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendVideoMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := NewVideoMessage(appId, "2cnEmOZzgSRcuadWzBrpm2bZqrqpWXa84yg7X42lasCiMVZf4H01XBi98-67KuCUV", "", 0)
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendFileMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := NewFileMessage(appId, "2aFk9yf0h4OfbDJt4q24LweZDxSfHXMGffGHetbj4hy3BbbURNTG59XwjFe7mHnJ9", 0)
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendTextCardMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := TextCardMessage{
		TextCard: struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			BtnTxt      string `json:"btntxt"`
		}{
			Title:       "测试卡片消息",
			Description: "测试卡片消息描述",
			URL:         "https://www.qq.com",
			BtnTxt:      "点击进入",
		},
	}
	m.MsgType = "textcard"
	m.AgentId = appId
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendNewsMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := NewsMessage{
		News: struct {
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
		},
	}
	m.MsgType = "news"
	m.AgentId = appId
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendMpNewsMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	m := MpNewsMessage{
		MpNews: struct {
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
		},
	}
	m.MsgType = "mpnews"
	m.AgentId = appId
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestApp_SendMarkdownMessage(t *testing.T) {
	corp := New(os.Getenv("COPP_ID"))
	appId, _ := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	app := corp.WithApp(appId, os.Getenv("APP_SECRET"))
	content := "您的会议室已经预定，稍后会同步到`邮箱`  \n>**事项详情**  \n>事　项：<font color=\"info\">开会</font>  \n>组织者：@miglioguan  \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang  \n>  \n>会议室：<font color=\"info\">广州TIT 1楼 301</font>  \n>日　期：<font color=\"warning\">2018年5月18日</font>  \n>时　间：<font color=\"comment\">上午9:00-11:00</font>  \n>  \n>请准时参加会议。  \n>  \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"
	m := NewMarkdownMessage(appId, content)
	m.ToUser = os.Getenv("TEST_USERID")
	resp, err := app.SendMessage(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
