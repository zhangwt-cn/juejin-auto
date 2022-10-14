package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"juejin-auto/model"
	"juejin-auto/util"
	"log"
	"net/http"
)


func checkIn(config model.Config) (string, string) {
	// 签到接口  url
	var url = "https://api.juejin.cn/growth_api/v1/check_in"
	respone := juejinReq(util.POST, url, config.Cookie)
	var markdownMsg string
	var textMsg string
	if respone.ErrNo == 0 {
		dataMap := respone.Data.(map[string]interface{})
		markdownMsg = fmt.Sprintf("  \n  - \u2705 签到成功，\U0001f38a 获得 **%v** 矿石～", dataMap["incr_point"])
		textMsg = fmt.Sprintf("\n\u2705 签到成功，\U0001f38a 获得 **%v** 矿石～", dataMap["incr_point"])
	} else {
		markdownMsg = "  \n  - \u274E 签到失败！ \n - \u2B07\uFE0F 失败原因：  \n > " + respone.ErrMsg
		textMsg = "\n\u274E 签到失败！\n\u2B07\uFE0F 失败原因：  \n    " + respone.ErrMsg
	}
	return markdownMsg, textMsg
}

// 通知
func notice(config model.Config, markdownMsg, textMsg string){
	sendMsg := util.MARKDOWN_TEXT + markdownMsg
	sendTextMsg := util.TEXT_CONTENT + textMsg
	util.SendDingtalkMsg(config.DingtalkBotToken, sendMsg, util.MARKDOWN)
	util.SendServerChanMsg(config.ServerChanToken, sendMsg)
	util.SendTelegramMsg(config.TelegramBotToken, config.ChatId, sendTextMsg)
}

// 获取账户矿石信息
func oreTotal(config model.Config) (string, string){
	checkTotalUrl := "https://api.juejin.cn/growth_api/v1/get_cur_point?aid=2608&uuid=6897007117560350216&spider=0"
	respone := juejinReq(util.GET, checkTotalUrl, config.Cookie)
	var markdownMsg string
	var textMsg string
	if respone.ErrNo == 0 {
		markdownMsg = fmt.Sprintf("  \n  - \U0001f389 当前矿石：**%v**", respone.Data)
		textMsg = fmt.Sprintf("\n\U0001f389 当前矿石：**%v**", respone.Data)
	} else {
		markdownMsg = fmt.Sprintf("  \n  - \u274E 获取当前矿石数量失败  \n - \u2B07\uFE0F 失败原因：  \n    > %s", respone.ErrMsg)
		textMsg = fmt.Sprintf("\n\u274E 获取当前矿石数量失败  \n\u2B07\uFE0F 失败原因：  \n    %s", respone.ErrMsg)
	}
	return markdownMsg, textMsg
}



// 获取账户签到信息
func checkInTotal(config model.Config) (string, string){
	checkTotalUrl := "https://api.juejin.cn/growth_api/v1/get_counts?aid=2608&uuid=6897007117560350216&spider=0"
	respone := juejinReq(util.GET, checkTotalUrl, config.Cookie)
	var markdownMsg string
	var textMsg string
	if respone.ErrNo == 0 {
		dataMap := respone.Data.(map[string]interface{})
		cont_count := dataMap["cont_count"]
		sum_count := dataMap["sum_count"]
		markdownMsg = fmt.Sprintf("  \n  - \U0001f389 连续签到 **%v** 天, 累计签到 **%v** 天", cont_count, sum_count)
		textMsg = fmt.Sprintf("\n\U0001f389 连续签到 **%v** 天, 累计签到 **%v** 天", cont_count, sum_count)
	} else {
		markdownMsg = fmt.Sprintf("  \n  - \u274E 获取签到信息失败!  \n - \u2B07\uFE0F 失败原因：  \n    > %s", respone.ErrMsg)
		textMsg = fmt.Sprintf("\n\u274E 获取签到信息失败!  \n\u2B07\uFE0F 失败原因：  \n    %s", respone.ErrMsg)
	}
	return markdownMsg, textMsg
}


// 掘金免费抽奖
func draw(config model.Config) (string, string) {
	myLuck := "https://api.juejin.cn/growth_api/v1/lottery_config/get"
	luckResp := juejinReq(http.MethodGet, myLuck, config.Cookie)
	var markdownMsg string
	var textMsg string
	if luckResp.ErrNo != 0 {
		markdownMsg = "  \n  - \u274E 获取免费抽奖次数失败，\u2B07\uFE0F 失败原因：\n > " + luckResp.ErrMsg
		textMsg = "\n\u274E 获取免费抽奖次数失败，\u2B07\uFE0F 失败原因：\n    " + luckResp.ErrMsg
		return markdownMsg, textMsg
	}
	luckRespData := luckResp.Data.(map[string]interface{})
	if luckRespData["free_count"].(float64) == 0 {
		markdownMsg = "  \n  - 免费抽奖次数为0，取消抽奖！"
		textMsg = "\n 免费抽奖次数为0，取消抽奖！" 
		return markdownMsg, textMsg
	}
	// 抽奖
	drawUrl := "https://api.juejin.cn/growth_api/v1/lottery/draw?aid=2608&uuid=7127091096610342439&spider=0&_signature=_02B4Z6wo00901SN-INAAAIDAnDGPCfkubJEjeyRAACua5TBxodFUaIpmIcDTKg3cY548fHqdK8GQIn.kNdU0DQuzJEH9KCq-x7q0rArPTDXh9U22vmFvnMyP.VWOk9sJ5aUStkTzTQj4rBNp01"
	drawResp := juejinReq(http.MethodPost, drawUrl, config.Cookie)
	if drawResp.ErrNo == 0 {
		dataMap := drawResp.Data.(map[string]interface{})
		markdownMsg = fmt.Sprintf("  \n  -  \U0001f3b0免费抽奖成功， 抽中 %v , 获得 %v 幸运值, 当前 %v 幸运值 ", dataMap["lottery_name"], dataMap["draw_lucky_value"], dataMap["total_lucky_value"])
		textMsg = fmt.Sprintf("\n\U0001f3b0免费抽奖成功， 抽中 %v , 获得 %v 幸运值, 当前 %v 幸运值", dataMap["lottery_name"], dataMap["draw_lucky_value"], dataMap["total_lucky_value"])
	} else {
		markdownMsg = "  \n  - \u274E 免费抽奖失败，\u2B07\uFE0F 失败原因：\n > " + drawResp.ErrMsg
		textMsg = "\n\u274E 免费抽奖失败，\u2B07\uFE0F 失败原因：\n    " + drawResp.ErrMsg
	}
	return markdownMsg, textMsg
}



// 掘金接口请求
func juejinReq(method, url, cookie string) model.Resp {
	headMap := map[string]string{
		"cookie": cookie,
	}
	resp, err := util.Req(method, url, headMap)
	if err != nil {
		log.Fatal("掘金接口请求异常", err)
	}
	defer resp.Body.Close()
	var respone model.Resp
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(data))
	json.Unmarshal([]byte(string(data)), &respone)
	return respone
}

// msg 组装
func msg(config model.Config) (string, string){
	markdownMsg, textMsg := checkIn(config)
	oreMarkdownMsg, oreTextMsg := oreTotal(config)
	checkInTotalMarkdownMsg, checkInTotalTextMsg := checkInTotal(config)
	drawMsg, drawTextMsg := draw(config)
	return markdownMsg + oreMarkdownMsg + checkInTotalMarkdownMsg + drawMsg, textMsg + oreTextMsg + checkInTotalTextMsg + drawTextMsg
}

// 掘金任务
func Task(config model.Config){
	markdownMsg, textMsg := msg(config)
	notice(config, markdownMsg, textMsg)
}






