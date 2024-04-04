package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	http "net/http"
	"temperatureblanket/gettemperature/domain"
)

func GetTemperature(cityId int, token string) {
	cityIdString := fmt.Sprint(cityId)
	url := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/climate/temperature/locale/%s?token=%s", cityIdString, token)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	fmt.Print(resp.Body)
}

func GetId(country string, state string, city string, token string) []domain.IdInfo {
	url := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/locale/city?name=%s&state=%s&country=%s&token=%s",
		city,
		state,
		country,
		token,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var idfindo []domain.IdInfo

	err = json.Unmarshal(body, &idfindo)
	if err != nil {
		return nil
	}

	return idfindo
}

func PutCityId(cityId int, token string) {

	url := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api-manager/user-token/%s/locales", token)

	body := domain.LocaleId{
		Status:  "sucess",
		Locales: []int{cityId},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return
	}

	resp, err := http.NewRequest("PUT", url, bytes.NewReader(jsonData))
	if err != nil {
		return
	}

	fmt.Print(resp)
}
