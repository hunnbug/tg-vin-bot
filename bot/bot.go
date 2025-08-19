package bot

import (
	"fmt"
	"log"
	"os"
	"tgbot/models"
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

		newUser := models.CreateUser(msg.Sender.FirstName, msg.Sender.LastName, msg.Sender.Username)
		VIN := msg.Text

		if models.IsVIN(VIN) {
			return ctx.Send(fmt.Sprintf(
				"%s %s. Ваш ВИН валиден, информация по нему предоставлена ниже:",
				newUser.FirstName(),
				newUser.LastName(),
			))
		} else {
			return ctx.Send(fmt.Sprintf(
				"%s %s. Ваш ВИН невалиден, попробуйте другой.\nVIN должен состоять из 17 латинских букв и цифр, не содержать букв O, I, Q",
				newUser.FirstName(),
				newUser.LastName(),
			))
		}
	})

	//19XFB2650DE800899

}
