package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

var templates = template.Must(template.ParseFiles("edit.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body1 []byte
	Body2 []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	//return ioutil.WriteFile(filename, p.Body, 0600)

	f, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	if p.Body1 != nil {
		if _, err := f.WriteString(string(p.Body1) + "\n"); err != nil {
			log.Println(err)
		}
	}
	if p.Body2 != nil {
		if _, err := f.WriteString(string(p.Body2) + "\n"); err != nil {
			log.Println(err)
		}
	}

	return nil
}
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	filename2 := title + "2.txt"
	body1, err := ioutil.ReadFile(filename)

	body2, err := ioutil.ReadFile(filename2)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body1: body1, Body2: body2}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	body2 := r.FormValue("body2")
	p := &Page{Title: title, Body1: []byte(body), Body2: []byte(body2)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/edit/test", http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
