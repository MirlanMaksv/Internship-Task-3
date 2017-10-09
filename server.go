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

	//link := "/getmusic https://www.youtube.com/watch?v=nfs8NYg7yQM"
	//videoId := "nfs8NYg7yQM"
	//prefix := "321"
	//model.PrepareAudio(link, types.ShortMessage{}, prefix)
}
