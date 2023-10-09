package config

import (
	"fmt"
	"time"
	"unitclock/ml-logger/server/global"

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

	global.R.AccessToken = viper.GetString("access_token")
	global.R.ExperimentId = viper.GetString("experiment_id")
	global.R.ProjectId = viper.GetString("project_id")
	global.R.StartAt = time.Now().Format("2006-01-02 15:04:05")

	Config = viper.GetViper()
}
