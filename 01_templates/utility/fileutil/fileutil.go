package fileutil

import (
	"log"
	"os"
	"strings"
)

func OutF(path, tn string) *os.File {
	if len(path) < 1 {
		return os.Stdout
	}

	sx := strings.Split(tn, ".")
	if len(sx) != 2 {
		log.Panicf("Expected %s to split into 2 but got: %d", tn, len(sx))
	}
	fn := strings.Join([]string{path, sx[0]}, "")
	fn = strings.Join([]string{fn, "html"}, ".")
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Panic("Unable to open: ", fn, " ", err)
	}
	return f
}

func CloseF(f *os.File) {
	if f == os.Stdout {
		return
	}

	log.Printf("Closing: %p, %v, %s", f, *f, f.Name())
	f.Close()
}
