package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	//"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	log.Println("Got connection...")
	fmt.Fprintf(conn, "I dialed you !..%d\n", 0)
	handle(conn)
}

func handle(conn net.Conn) {
	
	log.Println("Handling..")
	scn := bufio.NewScanner(conn)
	log.Println("Got scanner...")
	for scn.Scan() {
		txt := scn.Text()
		log.Println(txt)
		_, err := fmt.Fprintln(conn, "->Client says You Said: ", txt)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Exiting handle..", scn.Err())
}