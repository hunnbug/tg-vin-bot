package main

import (
	"log"
	"os"
	"time"

	env "github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {

	err := env.Load()
	if err != nil {
		log.Println("Произошла ошибка при загрузке env файла:", err)
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: time.Second * 5},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Println("Не удалось создать бота:", err)
	}

	bot.Handle("/hello", func(ctx tele.Context) error {
		return ctx.Send("Привет!")
	})

	log.Println("Бот начал работу")
	bot.Start()
}
