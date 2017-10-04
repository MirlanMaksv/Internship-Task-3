package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	page, err := loadPage("index")
	if err == nil {
		w.Write(page.Body)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", handler)
	http.ListenAndServe(":8080", nil)
}
