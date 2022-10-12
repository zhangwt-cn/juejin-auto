package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const(
	MARKDOWN = "markdown"
	TEXT     = "text"
	MARKDOWN_TEXT = "# auto task by juejin notice"
	TEXT_CONTENT = "auto task by juejin notice"
	DINGTALK_JSON = `{
		"msgtype": "%s",
		"markdown": {
			"title":"auto task by juejin notice",
			"text": "%s"
		},
		"text": {"content":"%s"}
		}`
	SERVER_CHAN_JSON = `{
		"title": "auto task by juejin notice",
		"desp": "%s",
		"channel": "9"
	}`
)


// 发送消息
// token bot token
// 发送消息
func SendDingtalkMsg(token, msg, msgType string){
	if token == "" {
		log.Println("未配置 dingtalk bot token 取消dingtalk推送!")
		return
	}

	botUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	data := fmt.Sprintf(DINGTALK_JSON, msgType, msg, msg)
	resp, _ := http.Post(botUrl, "application/json", strings.NewReader(data))
	byteArr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("dingtalk return：%s" ,string(byteArr))
}

// 发送 server chan 消息通知
func SendServerChanMsg(token, msg string){
	if token == "" {
		log.Println("未配置 server chan token 取消server chan推送!")
		return
	}
	serverChanUrl := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", token)
	data := fmt.Sprintf(SERVER_CHAN_JSON, strings.ReplaceAll(msg, "\n", "\\n")) 
	log.Printf("server chan req: %s" , data)
	resp, _ := http.Post(serverChanUrl, "application/json", strings.NewReader(data))
	byteArr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("server chan return：%s", string(byteArr))
}