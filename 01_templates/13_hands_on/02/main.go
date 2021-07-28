package main

import (
	"log"
	"os"
	"text/template"

	"github.com/sabubhatia/golang-webdev/01_templates/utility/fileutil"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Region string
	Hotels []hotel
}

type regions []region

const (
	Southern = "Southern"
	Central  = "Central"
	Northern = "Northern"
)

var tpl *template.Template

func init() {
	if len(os.Args) < 2 {
		log.Fatalf("At least 2 arguments expected. Got: %d\n", len(os.Args))
	}
	tpl = template.Must(template.ParseGlob(os.Args[1]))
}

func main() {
	r := regions{
		{
			Region: Southern,
			Hotels: []hotel{
				{
					Name:    "Hyatt",
					Address: "San Jose Sea world",
					City:    "San Jose",
					Zip:     "32878",
				},
				{
					Name:    "Four Seasons",
					Address: "San Jose Pier",
					City:    "San Jose",
					Zip:     "32625",
				},
			},
		},
		{
			Region: Central,
			Hotels: []hotel{
				{
					Name:    "Mandarin Oriental",
					Address: "Hollywood",
					City:    "Los Angeles",
					Zip:     "22178",
				},
			},
		},
		{
			Region: Northern,
			Hotels: []hotel{
				{
					Name:    "Langham Suites",
					Address: "Wine Yards Heaven",
					City:    "Napa",
					Zip:     "254378",
				},
				{
					Name:    "Westin",
					Address: "Napa Gardens",
					City:    "Napa",
					Zip:     "23498",
				},
			},
		},
	}

	f := fileutil.OutF("./tmp/", "HotelListings")
	defer fileutil.CloseF(f)
	err := tpl.ExecuteTemplate(f, "HotelListings", r)
	if err != nil {
		log.Panicf("Unable to execute template: %s\n", err.Error())
	}
}
