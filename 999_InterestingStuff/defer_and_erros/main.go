package main

import (
	"fmt"
	"log"
)

func cleanUp() error {
	return fmt.Errorf("**ERROR: cleanUp() error")
}

func doSomething() error {
	return fmt.Errorf("**ERROR: doSomething() error")
}
func getMessageBug() (string, error) {
	defer cleanUp()

	return "This is buggy. Where is the cleanUp() error?", nil
}

func getMessageBug2() (string, error) {
	var err error
	s := "Ok"

	fmt.Println(&err, err)
	defer func() {
		err = cleanUp()
		s = "This too is buggy"
		fmt.Println(&err, err)
	}()

	return s, err
}

func main() {
	/*	msg, err := getMessageBug()
		if err != nil {
			log.Println(err)
		}
		log.Println(msg)
	*/
	msg, err := getMessageBug2()
	if err != nil {
		log.Println(err)
	}
	log.Println(msg)
}
