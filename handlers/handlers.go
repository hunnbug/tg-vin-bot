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

		err := ctx.Send("Посылаем запрос")
		if err != nil {
			return err
		}

		messages, err := infrastructure.OSAGORequest(VIN)
		if err != nil {
			log.Println("при отправке запроса на получение ОСАГО произошла ошибка:", err)
			return ctx.Send("произошла ошибка сервера, попробуйте снова в другой раз")
		}

		err = ctx.Send(fmt.Sprintf("%s %s. Информация найдена:\n", newUser.FirstName(), newUser.LastName()))
		if err != nil {
			return err
		}

		for _, message := range messages {
			err := ctx.Send(message)
			if err != nil {
				return err
			}
		}

		return ctx.Send("Вы можете продолжить работу с ботом, для продолжения работы отправьте VIN номер авто")
	} else {
		return ctx.Send(fmt.Sprintf(
			"%s %s. Ваш ВИН невалиден.\nVIN должен состоять из 17 латинских букв и цифр, не содержать букв O, I, Q",
			newUser.FirstName(),
			newUser.LastName(),
		))
	}
}
