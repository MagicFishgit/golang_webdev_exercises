package main

import (
	"log"
	"os"
	"text/template"
)

type menuItem struct {
	Name, Description, Price string
}

type menuItems []menuItem

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	mItem := menuItems{
		menuItem{
			Name:        "MagicFish Surprise",
			Description: "Tastes a bit fishy and a lot of Magicy",
			Price:       "R300",
		},
		menuItem{
			Name:        "Saaas Surprise",
			Description: "Tastes friendy but with some salt could make it even better",
			Price:       "Infinite",
		},
		menuItem{
			Name:        "Cherie Special",
			Description: "High quality but may have you sit down for 4 hours.",
			Price:       "R900",
		},
	}

	err := tpl.Execute(os.Stdout, mItem)
	if err != nil {
		log.Fatalln(err)
	}

}
