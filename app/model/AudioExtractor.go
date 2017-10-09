package model

import (
	"fmt"
	"mirlan.maksv/telegram-bot/app"
)

func ExtractAudio(filename string, prefix string, videoExt string) (string, bool) {
	filename = prefix + filename
	command := fmt.Sprintf(app.Commands.FFMPEG, filename+"."+videoExt, filename)
	_, ok := Execute(command)
	if !ok {
		return "", false
	}
	return filename, true
}
