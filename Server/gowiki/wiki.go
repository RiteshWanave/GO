package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func save (p *Page) error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body,0600)
}

func load (title string) (*Page, error) {
	filename := title + ".txt"
	body, error := os.ReadFile(filename)
	p := &Page{Title: title, Body: body}
	if error != nil {
		return nil, error
	}
	return p, nil 
}

func viewHandler (w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, error := load(title)
	if error != nil {
		fmt.Println(http.StatusFound)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func edithandler (w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, error := load(title)
	if error != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func saveHandler (w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	e := save (&Page{Title: title, Body: []byte(body)})
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(http.StatusFound)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main () {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	save(p1)
	p2, _ := load("TestPage")
	fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", edithandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
