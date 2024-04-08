package main

import (
	"temperatureblanket/getweather/api"
	"temperatureblanket/getweather/configuration"
)

func main() {
	var config configuration.Configuration

	err := configuration.GetConfig(config)
	if err != nil {
		return
	}

	idInfo := api.GetId("BR", "SP", "São Paulo", config.Token)

	api.RegisterCityId(idInfo[0].Id, config.Token)

	//api.GetRegisteredCityByToken(token)

	//api.GetCurrentWeather(idInfo[0].Id, token)

}
