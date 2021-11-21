package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	rcvd := "null"

	var a []string

	if err := json.Unmarshal([]byte(rcvd), &a); err != nil {
		log.Fatal(err)
	}

	fmt.Println("a: ", a, "len: ", len(a), "cap: ", cap(a))
	bs, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bs: ", bs, string(bs))

	var ip *int
	if err := json.Unmarshal([]byte(rcvd), &ip); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ip: %v %T\n", ip, ip)
}
