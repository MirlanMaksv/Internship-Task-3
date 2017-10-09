package model

import (
	"fmt"
	"os/exec"
	"strings"
	"net/url"
	"encoding/json"
	"net/http"
	"bytes"
	"mirlan.maksv/telegram-bot/app/model/types"
	"mirlan.maksv/telegram-bot/app/telegram"
	"mirlan.maksv/telegram-bot/app"
)

func PrepareAudio(link string, msg types.ShortMessage, prefix string) {
	id, ok := ExtractId(link)
	if !ok {
		msg.Text = telegram.Messages.SmtWentWrong
		Send(telegram.Methods.SendMessage, msg)
		return
	}
	res, ok := DownloadVideo(id, prefix)
	ext := GetExtension(res, id)
	filename, ok := ExtractAudio(id, prefix, ext)
	if !ok {
		msg.Text = telegram.Messages.SmtWentWrong
		Send(telegram.Methods.SendMessage, msg)
		return
	}
	msg.Text = telegram.Messages.Uploading
	Send(telegram.Methods.SendMessage, msg)

	audio := types.Audio{Chat_id: msg.Chat_id, Audio: app.Config.AppUrl + "/get?link=" + filename}
	Send(telegram.Methods.SendAudio, audio)
}

// Used to get extension of a video file after it's downloaded
func GetExtension(data string, filename string) string {
	arr := strings.Split(data, "\n")
	for _, e := range arr {
		if strings.Contains(e, filename+".") {
			ext := strings.Split(e, filename+".")[1]
			for i := 0; i < len(ext); i++ {
				if ext[i] == ' ' {
					ext = ext[:i]
					break
				}
			}
			return ext
		}
	}
	return ""
}

func ExtractId(link string) (string, bool) {
	arr := strings.Fields(link)
	if len(arr) <= 1 {
		return "", false
	}
	URL := arr[1]
	urlParsed, err := url.Parse(URL)
	if err != nil {
		return "", false
	}
	params, err := url.ParseQuery(urlParsed.RawQuery)
	if err != nil {
		return "", false
	}
	value := params.Get("v")
	if len(value) == 0 {
		return "", false
	}
	return value, true
}

func Execute(command string) (string, bool) {
	split := strings.Fields(command)
	out, err := exec.Command(split[0], split[1:]...).Output()
	if err != nil {
		return "", false
	}
	return fmt.Sprintf("%s", out), true
}

func Send(method string, msg interface{}) {
	reqUrl := fmt.Sprintf(telegram.Url, telegram.BotToken, method)
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	http.Post(reqUrl, "application/json", bytes.NewBuffer(data))
}
