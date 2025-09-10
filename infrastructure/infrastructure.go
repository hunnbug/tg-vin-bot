package infrastructure

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"tgbot/models"
)

func OSAGORequest(vin string) (models.OSAGO, error) {

	baseUrl := "https://api-cloud.ru/api/rsa.php"

	params := url.Values{}
	params.Add("type", "osago")
	params.Add("vin", vin)
	params.Add("token", os.Getenv("API_TOKEN"))

	fullUrl := baseUrl + "?" + params.Encode()

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Произошла ошибка при отправке запроса на апи:", err)
		return models.OSAGO{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Произошла ошибка при чтении ответа от сервера:", err)
		return models.OSAGO{}, err
	}

	log.Println("получен ответ:", resp.StatusCode)

	data := make(map[string]any)

	_ = json.Unmarshal(body, &data)
	log.Println("анмаршал:", data)
	var osago models.OsagoResponse
	err = json.Unmarshal(body, &osago)
	if err != nil {
		log.Println("Произошла ошибка при анмаршаллинге ответа от апи:", err)
		return models.OSAGO{}, err
	}

	log.Println("Ответ от АПИ ОСАГО получен успешно")
	log.Printf("Анмаршал в структуру: %v\n\n%v", osago, osago.Rez[0])

	return osago.Rez[0], nil

}
