package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pictures struct {
	State bool
	Name  string
	Pics  []string
}

type private struct {
	state bool `json:"State"`
	name  string `json:"Name"`
	pics  []string
}

func main() {
	p := pictures{}
	bs, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p: ", p, ",", "bs: ", bs, string(bs))

	if err := json.Unmarshal(bs, &p); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarshalled: ", "p: ", p, ",", "bs: ", bs, string(bs))

	p = pictures {
		State: true,
		Name: "Picassos",
		Pics: []string {
			"One.jpg",
			"Two.jpg",
			"Three.jpg",
		},
	}

	bs, err = json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("p: ", p, "\n", "bs: ", bs, string(bs))

	var pr private
	if err := json.Unmarshal(bs, &pr); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unmarshalled: ", "pr: ", pr, "\n", "bs: ", bs, string(bs))

}
