package service

import "temperatureblanket/getweather/configuration"

type TemperatureData struct {
	Config configuration.Configuration
}

// GetWeather get info about the weather today
func (t *TemperatureData) GetWeather() {

}

// SendWeather sends the weather information by kafka
func (t *TemperatureData) SendWeather() {

}
