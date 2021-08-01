package rot13

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"time"
)

func TestRot13(t *testing.T) {
	bs := rot13([]byte("test"))
	wbs := []byte("grfg")
	if bytes.Compare(bs, wbs) != 0 {
		t.Errorf("Want %s, Got: %s", string(bs), string(wbs))
	}
	bs = rot13(wbs)
	wbs = []byte("test")
	if bytes.Compare(bs, wbs) != 0 {
		t.Errorf("Want: %s, test, Got: %s", string(wbs), string(bs))
	}
}

func TestServer(t *testing.T) {
	startServer()
}

func startServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()
	log.Println("Listening..")
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 30))
	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		txt := strings.ToLower(scn.Text())
		fmt.Fprintln(conn, "You say: ", txt)
		bs := rot13([]byte(txt))
		fmt.Fprintln(conn, "I Say: ", string(bs))
	}
	log.Println("Exiting handle()...")
}

func rot13(bs []byte) []byte {
	// We expect only lower case letters for now. 97 -122
	// Use ceaser cipher to rotate by 13.
	if len(bs) < 1 {
		return []byte{}
	}

	r13 := make([]byte, len(bs))
	for i, v := range bs {

		switch {
		case v < 97 || v > 122:
			log.Panicf("Expected charset is [a..z]. Got: %s\n", string(v))
		case v <= 109:
			r13[i] = v + 13
		default:
			r13[i] = v - 13
		}
	}

	return r13
}
