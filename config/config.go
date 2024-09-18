package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"github.com/labstack/gommon/log"
)

var conf = Config{}

func Get() Config {
	return conf
}

func Load() {
	defer setConfig()

	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")
	validConsul := consulURL != "" && consulPath != ""

	if validConsul {
		readFromConsul(consulURL, consulPath)
		return
	}

	readFromFile()
}

func readFromConsul(url, path string) {
	_ = viper.AddRemoteProvider("consul", url, path)
	viper.SetConfigType("json")

	if err := viper.ReadRemoteConfig(); err != nil {
		log.Error(err)
	}
}

func readFromFile() {
	viper.SetConfigFile("./config.local.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Info(err)
	}
}

func setConfig() {
	if err := viper.Unmarshal(&conf); err != nil {
		log.Error(err)
	}
}
