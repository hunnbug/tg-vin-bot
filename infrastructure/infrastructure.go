package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"tgbot/models"
)

const notFound string = "Не найдено"
const badTerm string = "Период использования транспортного средства равен сроку страхования. Дата, на которую запрошены сведения, не входит в период использования транспортного средства"

func OSAGORequest(vin string) ([]string, error) {

	baseUrl := "https://api-cloud.ru/api/rsa.php"

	params := url.Values{}
	params.Add("type", "osago")
	params.Add("vin", vin)
	params.Add("token", os.Getenv("API_TOKEN"))

	fullUrl := baseUrl + "?" + params.Encode()

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Произошла ошибка при отправке запроса на апи:", err)
		return []string{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Произошла ошибка при чтении ответа от сервера:", err)
		return []string{}, err
	}

	log.Println("получен ответ:", resp.StatusCode)

	var osagoResponse models.OsagoResponse
	err = json.Unmarshal(body, &osagoResponse)
	if err != nil {
		log.Println("Произошла ошибка при анмаршаллинге ответа от апи:", err)
		return []string{}, err
	}

	result := makeResult(osagoResponse)

	log.Println("Ответ от АПИ ОСАГО получен успешно")

	return result, nil

}

func makeResult(osagoResponse models.OsagoResponse) []string {

	result := make([]string, 4, 4)
	osago := osagoResponse.Rez[0]

	result[0] = fmt.Sprintf(`
	📋 Информация о полисе ОСАГО:
		• Серия и номер: %s %s
		• Страховая компания: %s
		• Статус: %s`,
		osago.Seria,
		osago.Nomer,
		osago.OrgOsago,
		osago.Status)

	if osago.Term == badTerm {
		result[1] = fmt.Sprintf(`
	📅 Сроки действия:
		• Срок действия полиса ОСАГО на данный момент окончен`)
	} else {
		result[1] = fmt.Sprintf(`
	📅 Сроки действия:
		• Период использования: %s
		• Начало использования: %s
		• Окончание использования: %s
		• Действие договора: с %s по %s`,
			osago.Term,
			osago.TermStart,
			osago.TermStop,
			osago.StartPolis,
			osago.StopPolis,
		)
	}

	if osago.RegNum == "" {
		result[2] = fmt.Sprintf(`
	🚗 Информация об автомобиле:
		• Марка и модель: %s
		• Гос. номер: %s`,
			osago.BrandModel,
			"Не найден")
	} else {
		result[2] = fmt.Sprintf(`
	🚗 Информация об автомобиле:
		• Марка и модель: %s
		• Гос. номер: %s`,
			osago.BrandModel,
			osago.RegNum)
	}

	result[3] = fmt.Sprintf("🌍 Расширение на Беларусь: %s", osago.DopBelarus)

	return result

}
