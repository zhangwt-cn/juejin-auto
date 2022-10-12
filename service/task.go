package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"juejin-auto/model"
	"juejin-auto/util"
	"log"
)


func checkIn(config model.Config) string {
	// 签到接口  url
	var url = "https://api.juejin.cn/growth_api/v1/check_in"
	respone := juejinReq(util.POST, url, config.Cookie)
	var msg string
	if respone.ErrNo == 0 {
		dataMap := respone.Data.(map[string]interface{})
		msg = fmt.Sprintf("  \n  - \u2705 签到成功，\U0001f38a 获得 **%v** 矿石～", dataMap["incr_point"])
	} else {
		msg = "  \n  - \u274E 签到失败！ \n - \u2B07\uFE0F 失败原因：  \n > " + respone.ErrMsg
	}
	return msg
}

// 通知
func notice(token, msg string){
	util.SendMsg(token, util.MARKDOWN_TEXT + msg, util.MARKDOWN)
}

// 获取账户矿石信息
func oreTotal(config model.Config) string{
	checkTotalUrl := "https://api.juejin.cn/growth_api/v1/get_cur_point?aid=2608&uuid=6897007117560350216&spider=0"
	respone := juejinReq(util.GET, checkTotalUrl, config.Cookie)
	var msg string
	if respone.ErrNo == 0 {
		msg = fmt.Sprintf("  \n  - \U0001f389 当前矿石：**%v**", respone.Data)
	} else {
		msg = fmt.Sprintf("  \n  - \u274E 获取当前矿石数量失败  \n - \u2B07\uFE0F 失败原因：  \n > %s", respone.ErrMsg)
	}
	return msg
}



// 获取账户签到信息
func checkInTotal(config model.Config) string{
	checkTotalUrl := "https://api.juejin.cn/growth_api/v1/get_counts?aid=2608&uuid=6897007117560350216&spider=0"
	respone := juejinReq(util.GET, checkTotalUrl, config.Cookie)
	var msg string
	if respone.ErrNo == 0 {
		dataMap := respone.Data.(map[string]interface{})
		cont_count := dataMap["cont_count"]
		sum_count := dataMap["sum_count"]
		msg = fmt.Sprintf("  \n  - \U0001f389 连续签到 **%v** 天, 累计签到 **%v** 天", cont_count, sum_count)
	} else {
		msg = fmt.Sprintf("  \n  - \u274E 获取签到信息失败!  \n - \u2B07\uFE0F 失败原因：  \n > %s", respone.ErrMsg)
	}
	return msg
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
func msg(config model.Config) string{
	if config.Cookie == "" {
		return "未配置掘金 token"
	}
	return checkIn(config) + oreTotal(config) + checkInTotal(config)
}

// 掘金任务
func Task(config model.Config){
	notice(config.DingtalkBotToken, msg(config))
}






