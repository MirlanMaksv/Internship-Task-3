package util

import (
	"os"
	"encoding/json"
	"fmt"
)

var wd string

func ParseJson(filename string, destination interface{}) {
	file, _ := os.Open(filename)
	json.NewDecoder(file).Decode(&destination)
}

func RemoveFile(filename string) {
	err := os.Remove(GetWd() + "/temp/" + filename)
	if err != nil {
		fmt.Println("Couldn't remove a file", err)
	}
}

func GetWd() string {
	if wd == "" {
		wd = os.Getenv("GOPATH") + "/src/mirlan.maksv/telegram-bot"
	}
	return wd
}