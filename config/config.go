package config

import (
	"encoding/json"
	"io/ioutil"
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

func init() {
	file, err := ioutil.ReadFile("config/config_dev.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &Conf)
	if err != nil {
		panic(err)
	}
}
