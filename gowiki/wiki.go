package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fmt.Println("Inside savePage")
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	fmt.Println("Inside loadPage")
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Cannot load file")
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func saveHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside saveHandler")
	title := req.URL.Path[len("/save/"):]
	body := "Hi This is my first wiki page"
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, req, "/view/"+title, http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside viewHandler")
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	fmt.Println("Inside wiki server")
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/view/", viewHandler)

	http.ListenAndServe(":9000", nil)
}
