package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	//viper.SetConfigFile("/go/src/github.com/jacky-htg/shonet-frontend/config/config.json")
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("ReadInConfig: ", err)
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
