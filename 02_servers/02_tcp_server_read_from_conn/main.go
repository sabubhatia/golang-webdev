package main

import (
	"bufio"
	"log"
	"net"
)


func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Listening: ", li.Addr())
	for {
		con, err := li.Accept()
		if err != nil {
			log.Println("Unable to accept connection: ", err)
		}
		log.Println("Openned conenction: ", con.LocalAddr().Network(), con.RemoteAddr().Network())
		go handle(con)
	}
}

func handle(con net.Conn) {

	scn := bufio.NewScanner(con)
	for scn.Scan() {
		log.Println(scn.Text())
	}

	defer con.Close()
	log.Println("Code got here...")
}