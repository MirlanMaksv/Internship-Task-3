package main

import (
	"mirlan.maksv/telegram-bot/app"
	"net/http"
	"mirlan.maksv/telegram-bot/app/controller"
)

func main() {
	 http.HandleFunc("/bot", controller.BotHandler)
	 http.HandleFunc("/get", controller.UploadHandler)
	 http.ListenAndServe(app.Config.AppPort, nil)
}
