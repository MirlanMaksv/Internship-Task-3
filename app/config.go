package app

import (
	"mirlan.maksv/telegram-bot/app/util"
)

type Configuration struct {
	AppUrl     string `json:"app_url"`
	AppPort    string `json:"app_port"`
}

var Config Configuration = Configuration{}

func init() {
	util.ParseJson(util.GetWd() + "/config.json", &Config)
}
