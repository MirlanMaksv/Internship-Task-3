package main

import (
	"os"
)

var Root = getWd() + "/src/mirlan.maksv/task#2/"
var Resources = "resources/"
var Temp = Root + "temp/"

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

var FFMPEG = "ffmpeg -i " + Temp + "%s " + Temp + "%s.mp3 -b:a 256k"
var YDL_FIRTS = "youtube-dl -f bestaudio %s"

// YDL_SECOND is needed as second half because fmt.Sprintf() is formatting them too
var YDL_SECOND = " -o " + Temp + "%(id)s.%(ext)s"
