package app

import (
	"mirlan.maksv/telegram-bot/app/util"
)

type Configuration struct {
	AppUrl     string `json:"app_url"`
	AppPort    string `json:"app_port"`
}

type CommandLine struct {
	FFMPEG string `json:"ffmpeg"`
	YDL_FIRST string `json:"ydl_first"`
	YDL_SECOND string `json:"ydl_second"`
}

var Config Configuration = Configuration{}
var Commands CommandLine = CommandLine{}

func init() {
	util.ParseJson("config.json", &Config)
	util.ParseJson("commands.json", &Commands)
}
