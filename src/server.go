package main

import (
	"html/template"
	"net/http"
	"os/exec"
	"fmt"
	"strings"
)

func main() {
	//test()
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/get", mp3Handler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func mp3Handler(w http.ResponseWriter, r *http.Request, ) {
	urlParam := r.URL.Query().Get("link")
	if len(urlParam) > 0 {
		command := fmt.Sprintf(YDL_FIRTS, urlParam) + YDL_SECOND
		res, ok := execute(command)
		if !ok {
			renderTemplate(w, "index", nil)
			return
		}
		ext := getExtension(res, urlParam)

		command = fmt.Sprintf(FFMPEG, urlParam+"."+ext, urlParam)
		execute(command)
	}
	renderTemplate(w, "result", &Page{Result: "SUCCESSFULLY DOWNLOADED"})
}

func execute(command string) (string, bool) {
	split := strings.Fields(command)
	out, err := exec.Command(split[0], split[1:]...).Output()
	if err != nil {
		return "", false
	}
	return fmt.Sprintf("%s", out), true
}

func getExtension(data string, filename string) string {
	out := strings.Split(data, "\n")
	for _, e := range out {
		if strings.Contains(e, filename+".") {
			ext := strings.Split(e, filename+".")[1]
			ext = ext[0:strings.Index(ext, " ")]
			return ext
		}
	}
	return ""
}

func renderTemplate(w http.ResponseWriter, tmp string, p *Page) {
	t, _ := template.ParseFiles(Root + Resources + tmp + ".html")
	t.Execute(w, p)
}

type Page struct {
	Result string
}
