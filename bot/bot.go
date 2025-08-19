package bot

import (
	"fmt"
	"log"
	"os"
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

	bot.Handle("/hello", func(ctx telebot.Context) error {
		return ctx.Send("Привет!", menu)
	})

	bot.Handle("/start", func(ctx telebot.Context) error {
		err := ctx.Send("Привет! Это бот для проверки информации об автомобиле по ВИН номеру")
		if err != nil {
			log.Println(err)
		}
		err = ctx.Send("Отправьте ваш ВИН ниже. (ВИН номер не может содержать русских букв, а так же букв O, I, Q)")

		return err
	})

	bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		msg := ctx.Message()
		user := msg.Sender

		return ctx.Send(fmt.Sprintf(
			"Вы написали: %s\nВы: %s %s %s",
			msg.Text,
			user.FirstName,
			user.LastName,
			user.Username,
		))
	})

	//19XFB2650DE800899

}
