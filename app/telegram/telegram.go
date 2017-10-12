package telegram

import (
	"mirlan.maksv/telegram-bot/app/util"
)

type Telegram struct {
	Url      string     `json:"url"`
	BotToken string     `json:"bot_token"`
	Commands TgCommands `json:"commands"`
	Methods  TgMethods  `json:"methods"`
	Messages TgMessages `json:"messages"`
}

type TgCommands struct {
	Start    string `json:"start"`
	GetMusic string `json:"getMusic"`
}

type TgMethods struct {
	SendMessage string `json:"sendMessage"`
	SendAudio   string `json:"sendAudio"`
}

type TgMessages struct {
	Introduce      string `json:"introduce"`
	NoArgs         string `json:"noArgs"`
	Wait           string `json:"wait"`
	SmtWentWrong   string `json:"smtWentWrong"`
	CantUnderstand string `json:"cantUnderstand"`
	Uploading      string `json:"uploading"`
}

var tg Telegram = Telegram{}

var Url string
var BotToken string
var Commands TgCommands
var Methods TgMethods
var Messages TgMessages

func init() {
	util.ParseJson(util.GetWd()+"/telegram.json", &tg)

	// reinitialize to minimize chaining
	Url = tg.Url
	BotToken = tg.BotToken
	Commands = tg.Commands
	Methods = tg.Methods
	Messages = tg.Messages
}
