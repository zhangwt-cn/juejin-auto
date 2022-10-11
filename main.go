package main

import (
	"juejin-auto/model"
	"juejin-auto/service"
	"os"
)

func main() {
	config := model.Config{Cookie: os.Getenv("JUEJIN_COOKIE")}
	service.CheckIn(config)
}