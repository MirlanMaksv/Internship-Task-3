package model

import (
	"fmt"
	"os/exec"
	"strings"
	"net/url"
	tg "mirlan.maksv/telegram-bot/app/telegram"
	"errors"
	"mirlan.maksv/telegram-bot/app/util"
)

var ffmpeg = util.GetWd() + "/scripts/ffmpeg.sh"
var youtubeDl = util.GetWd() + "/scripts/youtube-dl.sh"

func PrepareAudio(link string, prefix string) (string, string, error) {
	id, ok := extractId(link)
	if !ok {
		return "", tg.Messages.SmtWentWrong, errors.New("Couldn't extract id from link " + link)
	}

	res, ok := execute(youtubeDl, id, util.GetWd() + "/temp", prefix)
	if !ok {
		return "", tg.Messages.SmtWentWrong, errors.New("Error while downloading video " + link)
	}

	ext := getExtension(res, id)
	filename := prefix + id
	videoFilename := filename + "." + ext

	_, ok = execute(ffmpeg, videoFilename, util.GetWd() + "/temp", filename)
	if !ok {
		return "", tg.Messages.SmtWentWrong, errors.New("Couldn't extract audio from video " + videoFilename)
	}

	// remove video when it's downloaded and audio is extracted
	go util.RemoveFile(prefix + id + "." + ext)
	return filename, tg.Messages.Uploading, nil
}

// Used to get extension of a video file after it's downloaded
func getExtension(data string, filename string) string {
	arr := strings.Split(data, "\n")
	for _, e := range arr {
		if strings.Contains(e, filename+".") {
			ext := strings.Split(e, filename+".")[1]
			for i := 0; i < len(ext); i++ {
				if string(ext[i]) == " " {
					ext = string(ext[:i])
					break
				}
			}
			return ext
		}
	}
	return ""
}

func extractId(command string) (string, bool) {
	arr := strings.Fields(command)
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

func execute(command string, args ...string) (string, bool) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", false
	}
	return fmt.Sprintf("%s", out), true
}