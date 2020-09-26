package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./resources/")))
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", serveDog)
	http.HandleFunc("/dog.jpg", serveDogFile)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("/resources/index.html")
	if err != nil {
		log.Fatalln(w, "Oops!")
	}

	err = tpl.ExecuteTemplate(w, "/resources/index.html", nil)
	if err != nil {
		log.Fatalln(w, "There was a problem retrieving the site", err)
	}
}

func serveDog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("/resources/servedog.gohtml")
	if err != nil {
		log.Fatalln(w, "Oops!")
	}

	err = tpl.ExecuteTemplate(w, "/resources/servedog.gohtml", nil)
	if err != nil {
		log.Fatalln(w, "There was a problem retrieving the file", http.ErrMissingFile)
	}
}

func serveDogFile(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
