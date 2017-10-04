package main

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := page.Title + ".html"
	return ioutil.WriteFile(filename, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := Root + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
