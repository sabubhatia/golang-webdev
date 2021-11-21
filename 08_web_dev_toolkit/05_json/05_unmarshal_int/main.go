package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var i int
	rcvd := `42`

	if err := json.Unmarshal([]byte(rcvd), &i); err != nil {
		log.Fatal(err)
	}

	fmt.Println(i)

	bs, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}

	var j int
	fmt.Println("Json: ", bs, string(bs))
	json.NewDecoder(bytes.NewReader(bs)).Decode(&j)
	fmt.Println("J: ", j)

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(j)
	fmt.Println("Buf: ", buf.Bytes(), buf.String())
}
