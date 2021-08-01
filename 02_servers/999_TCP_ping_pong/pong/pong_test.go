package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)


func TestPong(t *testing.T) {
	pong(t)
}


func pong(t *testing.T) {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	t.Log("Listening...", li.Addr())
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err, conn)
		}
		log.Println("Got connection...")
		go handle(conn)
	}
}


func handle(conn net.Conn) {

	conn.SetDeadline(time.Now().Add(time.Second * 30))
	defer conn.Close()
	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		txt := scn.Text()
		log.Println(txt)
		fmt.Fprintln(conn, "Pong...")
	}

	log.Println("Exiting handle()..")
}
