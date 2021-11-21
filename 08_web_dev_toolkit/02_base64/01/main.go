package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	encoder := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	sabu := "SABUS"

	fmt.Println("sabu:", len(sabu))
	fmt.Println("encoder:", len(encoder))
	encoding := base64.NewEncoding(encoder)
	b64 := encoding.EncodeToString([]byte(sabu))
	fmt.Println("Encode: ", encoding.EncodedLen(len(sabu)), b64)
	s, err := encoding.DecodeString(b64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decode: ", encoding.DecodedLen(len(b64)), string(s))

}
