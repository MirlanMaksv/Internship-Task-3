package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"mirlan.maksv/telegram-bot/app/telegram"
)

func Send(method string, msg interface{}) {
	reqUrl := fmt.Sprintf(telegram.Url, telegram.BotToken, method)
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	http.Post(reqUrl, "application/json", bytes.NewBuffer(data))
}
