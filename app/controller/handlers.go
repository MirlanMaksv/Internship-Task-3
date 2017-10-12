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
	tg "mirlan.maksv/telegram-bot/app/telegram"
	"mirlan.maksv/telegram-bot/app/model/api"
	"mirlan.maksv/telegram-bot/app"
	"mirlan.maksv/telegram-bot/app/util"
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
	// remove audio when it's uploaded
	go util.RemoveFile(filename)
}

func BotHandler(w http.ResponseWriter, r *http.Request) {
	var in types.Message
	var tgMessage = types.TgMessage{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	tgMessage.Chat_id = in.Message.Chat.Id
	command := in.Message.Text

	isCommandCorrect := detectCommand(&tgMessage, command)
	api.Send(tg.Methods.SendMessage, &tgMessage)

	if isCommandCorrect {
		prefix := strconv.FormatUint(in.Update_id, 10)
		filename, resultMsg, err := model.PrepareAudio(command, prefix)
		if err != nil {
			fmt.Println(err)
			return
		}
		tgMessage.Text = resultMsg
		api.Send(tg.Methods.SendMessage, tgMessage)

		audio := types.Audio{
			Chat_id: tgMessage.Chat_id,
			Audio:   app.Config.AppUrl + "/get?link=" + filename}
		api.Send(tg.Methods.SendAudio, audio)
	}
}

func detectCommand(tgMessage *types.TgMessage, command string) bool {
	length := len(tg.Commands.GetMusic)
	if command == tg.Commands.Start {
		tgMessage.Text = tg.Messages.Introduce
	} else if command == tg.Commands.GetMusic {
		tgMessage.Text = tg.Messages.NoArgs
	} else if len(command) >= length && command[:length] == tg.Commands.GetMusic {
		tgMessage.Text = tg.Messages.Wait
		return true
	} else {
		tgMessage.Text = tg.Messages.CantUnderstand
	}
	return false
}
