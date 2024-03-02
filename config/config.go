package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

var conf Configuration

func GetConfig() Configuration {
	err := gonfig.GetConf("config/config.json", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
