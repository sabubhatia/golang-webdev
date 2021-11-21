package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	rcvd := `true`

	var b bool
	if err := json.Unmarshal([]byte(rcvd), &b); err != nil {
		log.Fatal(err)
	}

	fmt.Println("b: ", b)
	bs, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bs: ", bs, string(bs)) 

	var be bool
	if json.NewDecoder(bytes.NewReader(bs)).Decode(&be) != nil {
		log.Fatal(err)
	}

	fmt.Println("be: ", be)
}
