package bot

import (
	"log"
	"os"
	"tgbot/handlers"
	"time"

	"gopkg.in/telebot.v4"
)

func Init() (*telebot.Bot, error) {

	settings := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: time.Second * 5},
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		log.Println("Произошла ошибка при инициализации бота:", err)
	}

	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

	InitHandlers(bot, menu)

	return bot, err

}

func InitHandlers(bot *telebot.Bot, menu *telebot.ReplyMarkup) {

	bot.Handle("/start", handlers.HandleStart)
	bot.Handle(telebot.OnText, handlers.HandleVINSend)

	//19XFB2650DE800899

}
