package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	viper.AddConfigPath("./")

	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Read Config Failed :", err.Error())
	}
	Config = viper.GetViper()
}
