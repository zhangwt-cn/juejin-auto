package main

import (
	"juejin-auto/model"
	"juejin-auto/service"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	//定义一个结构体变量
	var config model.Config 
	log.Println(os.Getenv("JUEJIN_COOKIE")) 

	// 读取yaml文件
	byteArr, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Print(err)
	}
	// yaml文件内容映射到结构体中
	err1:= yaml.Unmarshal(byteArr, &config)
	if err1!=nil{
		log.Println("error")
	}
	service.CheckIn(config)
}