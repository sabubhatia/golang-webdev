package filelist

import (
	"log"
	"os"
)

func FileList(fx []string) func() *os.File {

	// If no files are in fx return os.Stdout. Else keep returning the next file in the list till
	// the end OF LIST IS HIT.

	useStdout := len(fx) <= 0
	next := 0
	return func() *os.File {
		if useStdout {
			return os.Stdout
		}
		var err error
		f, err := os.OpenFile(fx[next], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal("Error during file open: ", err)
		}
		next += 1
		return f
	}
}
