package controller

import (
	"encoding/json"
	"os"
	"net/http"
	"fmt"
	"io"
	"strconv"
	"mirlan.maksv/telegram-bot/app/model/types"
	"mirlan.maksv/telegram-bot/app/model"
	"mirlan.maksv/telegram-bot/app/telegram"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, ) {
	filename := r.URL.Query().Get("link")
	filename = "temp/" + filename + ".mp3"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	var in types.Message
	var msg = types.ShortMessage{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	msg.Chat_id = in.Message.Chat.Id
	command := in.Message.Text
	length := len(telegram.Commands.GetMusic)

	if command == telegram.Commands.Start {
		msg.Text = telegram.Messages.Introduce
	} else if command == telegram.Commands.GetMusic {
		msg.Text = telegram.Messages.NoArgs
	} else if len(command) >= length && command[:length] == telegram.Commands.GetMusic {
		msg.Text = telegram.Messages.Wait
		prefix := strconv.FormatUint(in.Update_id, 10)
		go model.PrepareAudio(command, msg, prefix)
	} else {
		msg.Text = telegram.Messages.CantUnderstand
	}

	model.Send(telegram.Methods.SendMessage, &msg)
}
