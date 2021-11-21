package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type statusCode struct {
	Code    int    `json:"code"`
	Descrip string `json:"descrip"`
}

type statusCodes []statusCode

func main() {
	var rcvd = ""
	var data statusCodes

	if err := json.Unmarshal([]byte(rcvd), &data); err != nil {
		log.Fatal(err)
	}

	for i, d := range data {
		fmt.Printf("%d. %d: %s\n", i, d.Code, d.Descrip)
	}

	bs, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json is: ")
	fmt.Println(string(bs))
}
