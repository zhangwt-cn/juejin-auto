package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	MARKDOWN     = "markdown"
	TEXT         = "text"
	MarkdownText = "# auto task by juejin notice"
	TextContent  = "auto task by juejin notice"
	DingTalkJson = `{
		"msgtype": "%s",
		"markdown": {
			"title":"auto task by juejin notice",
			"text": "%s"
		},
		"text": {"content":"%s"}
		}`
	ServerChanJson = `{
		"title": "auto task by juejin notice",
		"desp": "%s",
		"channel": "9"
	}`
	TelegramJson = `{
		"token":"%s",
		"msgText":"%s",
		"chatId":%s
	}`
)

// SendDingTalkMsg 发送消息
func SendDingTalkMsg(token, msg, msgType string) {
	if token == "" {
		log.Println("未配置 dingtalk bot token 取消dingtalk推送!")
		return
	}

	botUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	data := fmt.Sprintf(DingTalkJson, msgType, msg, msg)
	resp, _ := http.Post(botUrl, "application/json", strings.NewReader(data))
	byteArr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("dingtalk return：%s", string(byteArr))
}

// SendServerChanMsg 发送 server chan 消息通知
func SendServerChanMsg(token, msg string) {
	if token == "" {
		log.Println("未配置 server chan token 取消server chan推送!")
		return
	}
	serverChanUrl := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", token)
	data := fmt.Sprintf(ServerChanJson, strings.ReplaceAll(msg, "\n", "\\n"))
	log.Printf("server chan req: %s", data)
	resp, _ := http.Post(serverChanUrl, "application/json", strings.NewReader(data))
	byteArr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("server chan return：%s", string(byteArr))
}

// SendTelegramMsg 发送 Telegram 消息通知
func SendTelegramMsg(token, chatId, msg string) {
	if token == "" || chatId == "" {
		log.Println("未配置 Telegram token chatId 取消 Telegram 消息通知推送!")
		return
	}
	log.Printf("Telegram send:%s", msg)
	tgUrl := "https://tg-msg.vercel.app/api/tgNotify"
	data := fmt.Sprintf(TelegramJson, token, strings.ReplaceAll(msg, "\n", "\\n"), chatId)
	resp, _ := http.Post(tgUrl, "application/json", strings.NewReader(data))
	byteArr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Telegram return：%s", string(byteArr))
}
