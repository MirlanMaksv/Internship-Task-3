package main

import (
	"fmt"
	"strings"
	"os/exec"
)

func getAudio(videoId string, prefix string) (string, bool) {
	command := fmt.Sprintf(YDL_FIRST, videoId, prefix) + YDL_SECOND
	res, ok := execute(command)
	if !ok {
		return "", false
	}
	ext := getExtension(res, videoId)
	filename := prefix + videoId
	command = fmt.Sprintf(FFMPEG, filename+"."+ext, filename)
	res, ok = execute(command)
	if !ok {
		return "", false
	}
	return filename, true
}

func getExtension(data string, filename string) string {
	out := strings.Split(data, "\n")
	for _, e := range out {
		if strings.Contains(e, filename+".") {
			ext := strings.Split(e, filename+".")[1]
			for i := 0; i< len(ext); i++ {
				if ext[i] == ' ' {
					ext = ext[: i]
					break
				}
			}
			return ext
		}
	}
	return ""
}

func execute(command string) (string, bool) {
	split := strings.Fields(command)
	out, err := exec.Command(split[0], split[1:]...).Output()
	if err != nil {
		return "", false
	}
	return fmt.Sprintf("%s", out), true
}
