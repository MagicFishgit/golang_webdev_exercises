package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You  have reached index.")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You  have reached dog.")
}

func me(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("me.gohtml")
	if err != nil {
		log.Fatalln("There was an error parsing the template.", err)
	}

	err = tpl.ExecuteTemplate(w, "me.gohtml", "Rudi Visagie")
	if err != nil {
		log.Fatalln("There was an error executing the template", err)
	}
}

func main() {
	//handler
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
