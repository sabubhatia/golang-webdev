package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)


func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	log.Println("Listening: ", li.Addr().Network())
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		
		log.Println("Connection acquired: ", conn.RemoteAddr())
		go handle(conn)
		
	}
}

func handle(conn net.Conn) {

	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 10))
	scn := bufio.NewScanner(conn)
	log.Println("Got scanner...")
	for scn.Scan() {
		txt := scn.Text()
		log.Println("->", txt)
		_, err := fmt.Fprintln(conn, "Server says, you said: ", txt)
		if err != nil {
			log.Fatal(err)
		}

	}
	log.Println("Exiting handle()..")
}