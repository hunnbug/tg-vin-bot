package main

import (
	"log"
	"tgbot/bot"
	"tgbot/config"
)

func main() {

	err := config.LoadEnv()
	if err != nil {
		log.Println("Произошла ошибка при загруке env файла:", err)
		return
	}

	b, err := bot.Init()
	if err != nil {
		log.Println("Произошла ошибка при создании бота:", err)
		return
	}

	// bot.InitHandlers(b)

	log.Println("Бот начал работу")
	b.Start()
}
