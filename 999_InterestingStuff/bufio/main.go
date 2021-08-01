package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	stringBufio()
}

func stringBufio() {
	// Create a scanner off a sring.

	s := `The quick brown fox
	jumps over the lazy dog`

	scn := bufio.NewScanner(strings.NewReader(s))
	scn.Split(bufio.ScanWords)
	for scn.Scan() {
		fmt.Println(scn.Text())
	}
}
