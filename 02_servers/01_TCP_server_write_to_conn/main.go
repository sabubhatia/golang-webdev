package main

import (
	"fmt"
	"io"
	"log"
	"net"
)


func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()
	log.Printf("Server started...Network: %s @Address: %s\n", li.Addr().Network(), li.Addr().String())
	for {
		con, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		
		log.Println("Connection opened.", con.LocalAddr(), con.RemoteAddr())
		io.WriteString(con, "\n Hi there this is your friendly tcp server !\n");
		fmt.Fprintln(con, "How is your day?")
		fmt.Fprintf(con, "%v\n", "Good day I hope !")
		con.Close()
	}
}