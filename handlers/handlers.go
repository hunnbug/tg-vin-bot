package handlers

import (
	"fmt"
	"log"
	"tgbot/infrastructure"
	"tgbot/models"

	"gopkg.in/telebot.v4"
)

func HandleStart(ctx telebot.Context) error {

	err := ctx.Send("Привет! Это бот для проверки информации об автомобиле по ВИН номеру")
	if err != nil {
		log.Println(err)
	}
	err = ctx.Send("Отправьте ваш ВИН ниже. (ВИН номер не может содержать русских букв, а так же латинских букв O, I, Q)")

	return err
}

// Поменять на актуальные значения
const (
	Lightweight = "Легковая"
	Gazel       = "Легковая газель"
	Taxi        = "Такси"
	Highweight  = "Грузовой"
)

func HandleVINSend(ctx telebot.Context) error {
	msg := ctx.Message()

	newUser := models.CreateUser(msg.Sender.FirstName, msg.Sender.LastName, msg.Sender.Username)
	VIN := msg.Text

	if models.IsVIN(VIN) {

		ctx.Send("Посылаем запрос")

		osago, err := infrastructure.OSAGORequest(VIN)
		if err != nil {
			log.Println("при отправке запроса на получение ОСАГО произошла ошибка:", err)
			return ctx.Send("произошла ошибка сервера, попробуйте снова в другой раз")
		}

		message := fmt.Sprintf(`%s %s. Ваш ВИН валиден, информация по нему предоставлена ниже:
	📋 Информация о полисе ОСАГО:
			• Серия и номер: %s %s
			• Страховая компания: %s
			• Статус: %s

	📅 Сроки действия:
			• Период использования: %s
			→ Начало использования: %s
			→ Окончание использования: %s
			• Действие договора: с %s по %s

	🚗 Информация об автомобиле:
			• Марка и модель: %s
			• Гос. номер: %s
			• VIN: %s

	🌍 Расширение на Беларусь: %s`,
			newUser.FirstName(),
			newUser.LastName(),
			osago.Seria,
			osago.Nomer,
			osago.OrgOsago,
			osago.Status,
			osago.Term,
			osago.TermStart,
			osago.TermStop,
			osago.StartPolis,
			osago.StopPolis,
			osago.BrandModel,
			osago.RegNum,
			osago.VIN,
			osago.DopBelarus,
		)
		return ctx.Send(message)
	} else {
		return ctx.Send(fmt.Sprintf(
			"%s %s. Ваш ВИН невалиден.\nVIN должен состоять из 17 латинских букв и цифр, не содержать букв O, I, Q",
			newUser.FirstName(),
			newUser.LastName(),
		))
	}
}
