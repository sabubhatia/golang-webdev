package ping

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestPing(t *testing.T) {
	ping()
}

func ping() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("Cannot dial...", err)
	}
	log.Println("Got connection...!!")
	fmt.Fprintf(conn, "Ping..\n")
	handle(conn)
}


func handle(conn net.Conn) {
	defer conn.Close()
	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		txt := scn.Text()
		log.Println(txt)
		fmt.Fprintln(conn, "Ping...!")
	}
	log.Println("Exiting handle()...")
}