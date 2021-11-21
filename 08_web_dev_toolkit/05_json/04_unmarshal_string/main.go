package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	rcvd := `"Pheng"`

	var name string
	if err := json.Unmarshal([]byte(rcvd), &name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
