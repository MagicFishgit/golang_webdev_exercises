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

type restaurant struct {
	Name, Location string
	Menu           menuItems
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	restrnt := restaurants{
		restaurant{
			Name:     "Restaurant 1",
			Location: "Location 2",
			Menu: menuItems{
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
			},
		},
		restaurant{
			Name:     "Restaurant 2",
			Location: "Location 2",
			Menu: menuItems{
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
			},
		},
	}
	err := tpl.Execute(os.Stdout, restrnt)
	if err != nil {
		log.Fatalln(err)
	}

}
