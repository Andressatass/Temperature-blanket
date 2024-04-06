package main

import (
	"temperatureblanket/gettemperature/api"
)

func main() {
	token := "123"

	idInfo := api.GetId("BR", "SP", "São Paulo", token)

	api.RegisterCityId(idInfo[0].Id, token)

	//api.GetRegisteredCityByToken(token)

	api.GetCurrentWeather(idInfo[0].Id, token)

}
