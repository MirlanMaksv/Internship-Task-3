package model

import (
	"fmt"
	"mirlan.maksv/telegram-bot/app"
)

func DownloadVideo(videoId string, prefix string) (string, bool) {
	command := fmt.Sprintf(app.Commands.YDL_FIRST, videoId, prefix) + app.Commands.YDL_SECOND
	res, ok := Execute(command)
	if !ok {
		return "", false
	}
	return res, true
}
