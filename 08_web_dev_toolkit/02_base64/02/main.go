package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	sabu := "SABUS"
	b64 := base64.StdEncoding.EncodeToString([]byte(sabu))
	fmt.Println("Encode: ", base64.StdEncoding.EncodedLen(len(sabu)), b64)

	s, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decode: ", base64.StdEncoding.DecodedLen(len(b64)), s)
}
