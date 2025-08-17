package bot

import (
	"os"
	"time"

	"gopkg.in/telebot.v4"
)

func Init() (*telebot.Bot, error) {

	settings := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: time.Second * 5},
	}

	return telebot.NewBot(settings)

}

func InitHandlers(bot *telebot.Bot) {
	bot.Handle("/hello", func(ctx telebot.Context) error {
		return ctx.Send("Привет!")
	})
}
