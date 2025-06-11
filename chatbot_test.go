package workwx

import (
	"os"
	"testing"
)

func TestChatBot_SendTextMessage(t *testing.T) {
	chatBot := NewChatBot(os.Getenv("ROBOT_WEBHOOK"))
	m := ChatBotTextMessage{}
	m.MsgType = "text"
	m.Text.Content = "这是一个机器人测试消息"
	m.Text.MentionedList = []string{"@all", os.Getenv("TEST_USERID")}
	err := chatBot.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChatBot_SendMarkdownMessage(t *testing.T) {
	chatBot := NewChatBot(os.Getenv("ROBOT_WEBHOOK"))
	content := "您的会议室已经预定，稍后会同步到`邮箱`  \n>**事项详情**  \n>事　项：<font color=\"info\">开会</font>  \n>组织者：@miglioguan  \n>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang  \n>  \n>会议室：<font color=\"info\">广州TIT 1楼 301</font>  \n>日　期：<font color=\"warning\">2018年5月18日</font>  \n>时　间：<font color=\"comment\">上午9:00-11:00</font>  \n>  \n>请准时参加会议。  \n>  \n>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"
	m := ChatBotMarkdownMessage{}
	m.MsgType = "markdown"
	m.Markdown.Content = content
	err := chatBot.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChatBot_SendNewsMessage(t *testing.T) {
	chatBot := NewChatBot(os.Getenv("ROBOT_WEBHOOK"))
	m := ChatBotNewsMessage{}
	m.MsgType = "news"
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
	err := chatBot.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChatBot_SendFileMessage(t *testing.T) {
	chatBot := NewChatBot(os.Getenv("ROBOT_WEBHOOK"))
	m := ChatBotFileMessage{}
	m.MsgType = "file"
	m.File.MediaId = "2aFk9yf0h4OfbDJt4q24LweZDxSfHXMGffGHetbj4hy3BbbURNTG59XwjFe7mHnJ9"
	err := chatBot.Send(m)
	if err != nil {
		t.Fatal(err)
	}
}
