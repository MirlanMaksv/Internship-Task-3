package main

import (
	"html/template"
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"net/url"
	"strings"
	"os"
	"io"
)

func main() {
	http.HandleFunc("/get", uploadHandler)
	http.HandleFunc("/bot", botHandler)
	http.ListenAndServe(":8443", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, ) {
	filename := r.URL.Query().Get("link")
	filename = Temp + filename + ".mp3"
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

func botHandler(w http.ResponseWriter, r *http.Request) {
	var in Message
	var out = ShortMessage{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	out.Chat_id = in.Message.Chat.Id
	command := in.Message.Text
	length := len(TG_GETMUSIC)
	if command == TG_START {
		out.Text = MSG_INTRO
	} else if command == TG_GETMUSIC {
		out.Text = MSG_NO_ARGS
	} else if len(command) >= length && command[:length] == TG_GETMUSIC {
		out.Text = MSG_WAIT
		prefix := fmt.Sprintf("%d", in.Update_id)
		go prepareAudio(command, out, prefix)
	} else {
		out.Text = MSG_CANT_GET_YOU
	}
	send(TG_SENDMESSAGE, &out)
}

func prepareAudio(command string, msg ShortMessage, prefix string) {
	id, ok := extractId(command)
	if !ok {
		msg.Text = MSG_SMT_WENT_WRONG
		send(TG_SENDMESSAGE, msg)
		return
	}
	filename, ok := getAudio(id, prefix)
	if !ok {
		msg.Text = MSG_SMT_WENT_WRONG
		send(TG_SENDMESSAGE, msg)
		return
	}
	msg.Text = MSG_UPLOADING
	send(TG_SENDMESSAGE, msg)

	audio := Audio{Chat_id: msg.Chat_id, Audio: APP_URL + "/get?link=" + filename}
	send(TG_SENDAUDIO, audio)
}

func extractId(command string) (string, bool) {
	URL := strings.Fields(command)[1]
	urlParsed, err := url.Parse(URL)
	if err != nil {
		return "", false
	}
	params, err := url.ParseQuery(urlParsed.RawQuery)
	if err != nil {
		return "", false
	}
	return params.Get("v"), true
}

func send(method string, msg interface{}) {
	reqUrl := fmt.Sprintf(TG_URL, TG_TOKEN, method)
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	http.Post(reqUrl, "application/json", bytes.NewBuffer(data))
}

func renderTemplate(w http.ResponseWriter, tmp string, p *Page) {
	t, _ := template.ParseFiles(Root + Resources + tmp + ".html")
	t.Execute(w, p)
}
