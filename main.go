package main

import (
	"juejin-auto/model"
	"juejin-auto/service"
	"os"
)

func main() {
	config := model.Config{
		Cookie: os.Getenv("JUEJIN_COOKIE"),
		DingtalkBotToken: os.Getenv("DINGTALK_BOT_TOKEN"),
		ServerChanToken: os.Getenv("SERVER_CHAN_TOKEN"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ChatId: os.Getenv("TELEGRAM_CHAT_ID"),
	}
	service.Task(config)
}