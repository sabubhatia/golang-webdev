package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
)

// hashed based message authentication code.

func getCode(s string) string {
	bxx := bytes.Join([][]byte{[]byte("This is my super secret key"), []byte(s)}, []byte{})
	h := hmac.New(sha256.New, bxx)
	bx := h.Sum(nil)
	fmt.Printf("%v, Len: %d, Str: %s\n", bx, len(bx), string(bx))
	return fmt.Sprintf("%X", bx)
}

func main() {
	s1 := "phengsabu@gmail.com"
	s2 := "sabusingh.bhatia@gmail.com"

	h1 := getCode(s1)
	h2 := getCode(s2)

	fmt.Printf("S: %s, h: %v, L:%d\n", s1, h1, len(h1))
	
	fmt.Printf("S: %s, h: %v, L:%d\n", s2, h2, len(h2))

	if hmac.Equal([]byte(h1), []byte(h2)) {
		log.Println("Hmacs are equal")
		return
	}

	log.Println("Hmacs NOT equal")
}
