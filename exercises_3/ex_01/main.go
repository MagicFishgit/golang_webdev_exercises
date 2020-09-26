package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You  have reached index.")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You  have reached dog.")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are Rudi Visagie")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me) //handler
	http.ListenAndServe(":8080", nil)
}
