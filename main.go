package main

import (
	"html/template"
	"net/http"
)

/*var tmpl1 = template.Must(template.ParseGlob("/web/web/* "))*/
var tmpl = template.Must(template.ParseFiles("index.html", "elements.html", "blog.html", "ticket.html", "contact.html", "about_us.html", "footer.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about_us.html", nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}
func ticket(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "ticket.html", nil)
}
func blog(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "blog.html", nil)
}
func elements(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "elements.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("web/web"))
	http.Handle("/web/css", http.StripPrefix("/web/css", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/about_us", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/ticket", ticket)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/elements", elements)
	http.ListenAndServe(":8181", nil)
}
