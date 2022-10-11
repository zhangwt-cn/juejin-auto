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
	JSON = `{
		"msgtype": "%s",
		"markdown": {
			"title":"auto task by juejin notice",
			"text": "%s"
		},
		"text": {"content":"%s"}
		}`
)


// 发送消息
// token bot token
// 发送消息
func SendMsg(token, msg, msgType string){
	botUrl := "https://oapi.dingtalk.com/robot/send?access_token=" + token
	data := fmt.Sprintf(JSON, msgType, msg, msg)
	resp, _ := http.Post(botUrl, "application/json", strings.NewReader(data))
	byteArr, _ := ioutil.ReadAll(resp.Body)
	log.Printf("dingtalk 接口返回：%s" ,string(byteArr))
}