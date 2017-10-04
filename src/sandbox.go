package main

import (
	"fmt"
)

func test() {
	command := fmt.Sprintf(YDL_FIRTS, "3gzqsmx1KGU") + YDL_SECOND
	res, ok := execute(command)
	if !ok {
		return
	}
	ext := getExtension(res, "3gzqsmx1KGU")

	command = fmt.Sprintf(FFMPEG, "3gzqsmx1KGU"+"."+ext, "3gzqsmx1KGU")
	execute(command)
}
