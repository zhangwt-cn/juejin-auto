package service

import (
	"encoding/json"
	"io/ioutil"
	"juejin-auto/model"
	"juejin-auto/util"
	"log"
)


func CheckIn(config model.Config) {
	// 签到接口  url
	var url = "https://api.juejin.cn/growth_api/v1/check_in"

	headMap := map[string]string{
		"cookie":config.Cookie,
	}
	resp, err := util.Post(url, headMap)
	if err != nil {
		log.Fatal("签到接口请求异常", err)
	}
	defer resp.Body.Close()

	var respone model.Resp
	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(string(data)), &respone)
	verify(respone)
}

// 校验请求结果
func verify (resp model.Resp) {
	if resp.ErrNo == 0 {
		log.Println("签到成功！")
	} else {
		log.Printf("签到失败，失败原因：%s \n", resp.ErrMsg)
	}
	// TODO 通知
}






