package configuration

import (
	"fmt"

	"temperatureblanket/gettemperature/app"

	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
)

type Configuration struct {
	Token string `env:"GETWEATHER_TOKEN"`
}

func GetConfig(config Configuration) error {
	e := enviper.New(viper.New())

	e.SetEnvPrefix(app.EnvPrefix)
	e.SetConfigFile(app.ConfigFilePath)
	e.SetConfigName(app.ConfigName)

	err := e.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("could not unmarshal config into the struct: %s", err)
	}

	return nil
}
