package main

import (
	"temperatureblanket/gettemperature/api"
	"temperatureblanket/gettemperature/configuration"
)

func main() {

	var config configuration.Configuration

	err := configuration.GetConfig(config)
	if err != nil {
		return
	}

	token := config.Token

	idInfo := api.GetId("BR", "SP", "São Paulo", token)

	api.RegisterCityId(idInfo[0].Id, token)

	//api.GetRegisteredCityByToken(token)

	//api.GetCurrentWeather(idInfo[0].Id, token)

}
