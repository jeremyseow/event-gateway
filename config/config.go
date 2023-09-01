package config

import (
	"github.com/spf13/viper"
)

const (
	// bootstrap paths
	configBsPath = "bootstrap/config_dev.json"
)

type Config struct {
	Env        string
	HostConfig `json:"hostConfig"`
}

type HostConfig struct {
	Hostname  string `json:"hostName"`
	GrpcPort  string
	DebugPort string
}

var Conf = &Config{}

func InitConfig() *viper.Viper {
	viper.SetConfigFile(configBsPath)
	viper.AutomaticEnv() // viper will automatically override values that it has read from config file with the values of the corresponding environment variables if they exist
	// err := viper.BindEnv("env", "ENV")
	// if err != nil {
	// 	panic(err)
	// }
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// fmt.Println(viper.GetString("env"))

	return viper.GetViper()
}
