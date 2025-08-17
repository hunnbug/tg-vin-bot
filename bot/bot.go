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

	InitHandlers(bot, menu)

	return bot, err

}

func InitHandlers(bot *telebot.Bot, menu *telebot.ReplyMarkup) {

	// btnOSAGO := menu.Text("📄 Проверить полис ОСАГО 📄")
	// btnOSGOP := menu.Text("📄 Проверить Полис ОСГОП 📄")
	// btnMKAD := menu.Text("🚗 Проверить пропуск на МКАД 🚗")
	// btnTO := menu.Text("💳 Проверить диагностическую карту ТО 💳")
	// btnManager := menu.Text("☎️ Связаться с менеджером ☎️")
	// btnMessager := menu.Text("🤳 Подписаться на обновления 🤳")
	// btnShtrafi := menu.Text("💲 Проверить штрафы 💲")

	// menu.Reply(
	// 	menu.Row(btnOSAGO),
	// 	menu.Row(btnOSGOP),
	// 	menu.Row(btnTO),
	// 	menu.Row(btnMKAD),
	// 	menu.Row(btnManager),
	// 	menu.Row(btnMessager),
	// 	menu.Row(btnShtrafi),
	// )

	// bot.Handle(&btnOSAGO, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на осаго!")
	// })
	// bot.Handle(&btnOSGOP, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на осгоп!")
	// })
	// bot.Handle(&btnTO, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на ТО!")
	// })
	// bot.Handle(&btnMKAD, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на мкад!")
	// })
	// bot.Handle(&btnManager, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на менеджера!")
	// })
	// bot.Handle(&btnMessager, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на подписку!")
	// })
	// bot.Handle(&btnShtrafi, func(ctx telebot.Context) error {
	// 	return ctx.Send("Ты нажал на штрафы!")
	// })

	bot.Handle("/hello", func(ctx telebot.Context) error {
		return ctx.Send("Привет!", menu)
	})

	bot.Handle("/start", func(ctx telebot.Context) error {
		err := ctx.Send("Привет! Это бот для проверки информации об автомобиле по ВИН номеру")
		if err != nil {
			log.Println(err)
		}
		err = ctx.Send("Отправьте ваш ВИН ниже:")

		return err
	})

	//19XFB2650DE800899

}
