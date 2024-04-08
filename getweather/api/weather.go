package api

import (
	"encoding/json"
	"fmt"
	"io"
	http "net/http"
	neturl "net/url"
	"strings"
	"temperatureblanket/getweather/domain"
)

func GetCurrentWeather(cityId int, token string) {
	cityIdString := fmt.Sprint(cityId)
	url := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/weather/locale/%s/current?token=%s", cityIdString, token)

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

func GetRegisteredCityByToken(token string) {
	url := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api-manager/user-token/%s/locales", token)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var resgisteredCitys domain.ResgisteredCitys

	err = json.Unmarshal(body, &resgisteredCitys)
	if err != nil {
		return
	}

	fmt.Print(resgisteredCitys)
}

func RegisterCityId(cityId int, token string) {

	url := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api-manager/user-token/%s/locales", token)

	formData := neturl.Values{}
	formData.Add("localeId[]", fmt.Sprint(cityId))

	req, err := http.NewRequest("PUT", url, io.NopCloser(strings.NewReader(formData.Encode())))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	fmt.Println(string(body))
}
