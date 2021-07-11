package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid nuumber of command line arguments: Expected 2 got ", len(os.Args))
	}

	arg0 := os.Args[0]
	arg1 := os.Args[1]
	fmt.Printf("ARGS: %v, %v\n", arg0, arg1)

	str := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang=en>
	<head>
	<meta charset="UTF-8">
	<title> Hello World !</title>
	</head>
	<body>
	<h1> %s </h1>
	</body>
	</html>
	`, arg1)

	nf, err := os.Create("./tmp/index.html")
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}
	defer nf.Close()
	_, err = io.Copy(nf, strings.NewReader(str))
	if err != nil {
		log.Fatal("Error copying file.", err)
	}
}