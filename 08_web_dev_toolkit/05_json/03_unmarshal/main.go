package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type city struct {
	Precision string  `json:"precision"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Address   string  `json:"Address"`
	City      string  `json:"City"`
	State     string  `json:"State"`
	Zip       string  `json:"Zip"`
	Country   string  `json:"Country"`
}

type cities []city

type cityshort struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	City      string  `json:"City"`	
}

type citiesshort []cityshort

func main() {
	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},
	{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	var c cities

	if err := json.Unmarshal([]byte(rcvd), &c); err != nil {
		log.Fatal(err)
	}

	for _, v := range c {
		fmt.Println(v)
	}

	var cs citiesshort
	if err := json.Unmarshal([]byte(rcvd), &cs); err != nil {
		log.Fatal(err)
	}

	for _, v := range cs {
		fmt.Printf("%+v\n", v)
	}
}
