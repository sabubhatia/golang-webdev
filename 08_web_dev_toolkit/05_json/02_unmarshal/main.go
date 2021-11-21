package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL    string `json:"Url"`
	Height int    `json:"Height"`
	Width  int    `json:"Width"`
}

type pic struct {
	Width     int       `json:"Width"`
	Height    int       `json:"Height"`
	Title     string    `json:"Title"`
	Thumbnail thumbnail `json:"Thumbnail"`
	Animated  bool      `json:"Animated"`
	IDs       []int     `json:"IDs"`
}

func main() {
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	var p pic
	if err := json.Unmarshal([]byte(rcvd), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", p)
	fmt.Println(p)
	for _, v := range p.IDs {
		fmt.Printf("ID: %d\n", v)
	}
	fmt.Printf("URL: %s\n", p.Thumbnail.URL)
}
