package handlers

import (
	"fmt"
	"log"
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
}
