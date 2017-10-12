package model

import (
	"testing"
	"sync"
)

func TestDownloadVideo(t *testing.T) {
	var wg sync.WaitGroup
	url := "https://www.youtube.com/watch?v="
	ids := []string{"nfs8NYg7yQM", "JGwWNGJdvx8", "kJQP7kiw5Fk"}
	for _, v := range ids {
		wg.Add(1)
		go func (command string) {
			_, _, err := PrepareAudio(command, "1")
			if err != nil {
				t.Error("Error: ", err)
			}
			wg.Done()
		}("/getmusic " + url + v)
	}
	wg.Wait()
}

func TestIncorrectIds(t *testing.T) {
	var wg sync.WaitGroup
	url := "https://www.youtube.com/watch?v="
	ids := []string{"somePrefix_nfs8NYg7yQM"}
	for _, v := range ids {
		wg.Add(1)
		go func (command string) {
			_, _, err := PrepareAudio(command, "1")
			if err == nil {
				t.Error("Error: ", err)
			}
			wg.Done()
		}("/getmusic " + url + v)
	}
	wg.Wait()
}

// Incorrect query parameter
func TestIncorrectUrl(t *testing.T) {
	var wg sync.WaitGroup
	url := "https://www.youtube.com/watch?param="
	ids := []string{"nfs8NYg7yQM"}
	for _, v := range ids {
		wg.Add(1)
		go func (command string) {
			_, _, err := PrepareAudio(command, "1")
			if err == nil {
				t.Error("Error: ", err)
			}
			wg.Done()
		}("/getmusic " + url + v)
	}
	wg.Wait()
}

// Incorrect query parameter
func TestIncorrectCommand(t *testing.T) {
	var wg sync.WaitGroup
	url := "https://www.youtube.com/watch?param="
	ids := []string{"nfs8NYg7yQM"}
	for _, v := range ids {
		wg.Add(1)
		go func (command string) {
			_, _, err := PrepareAudio(command, "1")
			if err == nil {
				t.Error("Error: ", err)
			}
			wg.Done()
		}(url + v)
	}
	wg.Wait()
}