package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	htl := hotels{
		hotel{
			Name:    "Hotel Casablanca",
			Address: "42 Rando Street",
			City:    "Los Angeles",
			Zip:     "5880",
			Region:  "High Street",
		},
		hotel{
			Name:    "Hotel MagicFish",
			Address: "40 Rando2 Street",
			City:    "Los Angeles",
			Zip:     "5880",
			Region:  "Down Town",
		},
	}

	err := tpl.Execute(os.Stdout, htl)
	if err != nil {
		log.Fatalln(err)
	}

}
