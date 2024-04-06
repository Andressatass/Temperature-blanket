package main

import (
	"temperatureblanket/gettemperature/api"
	"temperatureblanket/gettemperature/config"
)

func main() {

	config := config.Configuration{}

	token := config.Token

	idInfo := api.GetId("BR", "SP", "São Paulo", token)

	api.RegisterCityId(idInfo[0].Id, token)

	//api.GetRegisteredCityByToken(token)

	api.GetCurrentWeather(idInfo[0].Id, token)

}
