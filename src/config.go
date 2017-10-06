package main

import (
	"os"
)

var APP_URL = "https://b91c4e5c.ngrok.io"
var Root = getWd() + "/src/mirlan.maksv/task#2/"
var Resources = "resources/"
var Temp = Root + "temp/"

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

var FFMPEG = "ffmpeg -i " + Temp + "%s " + Temp + "%s.mp3 -b:a 256k"
var YDL_FIRST = "youtube-dl -f bestaudio %s -o " + Temp + "%s"

// YDL_SECOND is needed as second half because fmt.Sprintf() is formatting them too
var YDL_SECOND = "%(id)s.%(ext)s"


/** TELEGRAM STUFF **/
var TG_URL = "https://api.telegram.org/bot%s/%s"
var TG_TOKEN = "467191911:AAFF5LPexqvVr4YoVuAmU-K77gCrkA555tw"

// @musicExtractorBot Telegram bot commands
var TG_START = "/start"
var TG_GETMUSIC = "/getmusic"

// Telegram Bot Api methods
var TG_SENDMESSAGE = "sendMessage"
var TG_SENDAUDIO = "sendAudio"

var MSG_INTRO = "Hi I am music extractor bot. Provide me with youtube video url you will get back mp3 of that video"
var MSG_NO_ARGS = "Give me URL"
var MSG_WAIT = "Downloading from the source ..."
var MSG_SMT_WENT_WRONG = "Oops, something went wrong :("
var MSG_CANT_GET_YOU = "Sorry, I didn't get you"
var MSG_UPLOADING = "Uploading ..."