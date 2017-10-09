package util

import (
	"os"
	"encoding/json"
)

func ParseJson(filename string, destination interface{}) {
	file, _ := os.Open(filename)
	json.NewDecoder(file).Decode(&destination)
}
