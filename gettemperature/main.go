package main

import (
	"temperatureblanket/gettemperature/api"
)

func main() {
	token := "123"

	idInfo := api.GetId("BR", "SP", "São Paulo", token)

	api.PutCityId(idInfo[0].Id, token)

	api.GetTemperature(idInfo[0].Id, token)

}
