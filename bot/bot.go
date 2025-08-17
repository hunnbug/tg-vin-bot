package bot

import (
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
	// selector := &telebot.ReplyMarkup{}

	InitHandlers(bot, menu)

	return bot, err

}

func InitHandlers(bot *telebot.Bot, menu *telebot.ReplyMarkup) {

	btnOSAGO := menu.Text("📄 Проверить полис ОСАГО 📄")
	btnOSGOP := menu.Text("📄 Проверить Полис ОСГОП 📄")
	btnMKAD := menu.Text("🚗 Проверить пропуск на МКАД 🚗")
	btnTO := menu.Text("💳 Проверить диагностическую карту ТО 💳")
	btnManager := menu.Text("☎️ Связаться с менеджером ☎️")
	btnMessager := menu.Text("🤳 Подписаться на обновления 🤳")

	menu.Reply(
		menu.Row(btnOSAGO),
		menu.Row(btnOSGOP),
		menu.Row(btnTO),
		menu.Row(btnMKAD),
		menu.Row(btnManager),
		menu.Row(btnMessager),
	)

	bot.Handle("/hello", func(ctx telebot.Context) error {
		return ctx.Send("Привет!", menu)
	})

	bot.Handle(&btnOSAGO, func(ctx telebot.Context) error {
		return ctx.Send("Ты нажал на осаго!")
	})
	bot.Handle(&btnOSGOP, func(ctx telebot.Context) error {
		return ctx.Send("Ты нажал на осгоп!")
	})
	bot.Handle(&btnTO, func(ctx telebot.Context) error {
		return ctx.Send("Ты нажал на ТО!")
	})
	bot.Handle(&btnMKAD, func(ctx telebot.Context) error {
		return ctx.Send("Ты нажал на мкад!")
	})
	bot.Handle(&btnManager, func(ctx telebot.Context) error {
		return ctx.Send("Ты нажал на менеджера!")
	})
	bot.Handle(&btnMessager, func(ctx telebot.Context) error {
		return ctx.Send("Ты нажал на подписку!")
	})
}
