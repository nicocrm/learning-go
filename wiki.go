package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err == nil {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	} else {
		fmt.Fprintf(w, "There was an error: %v", err)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	// p1.Save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
}
