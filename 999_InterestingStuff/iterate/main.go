package main

import "fmt"

// some code to get my head around returning functions
// We will create an iterator here:
// (1) Pass a slice of strings to iterate over.
// (2) Iterator function returns an iterator that will return a function which when invoked will return the next string.
// (3) The iterator returns nil if have already iterated over each element else a function.

func makeItr(s []string) func() func() string {
	nxt := 0
	return func() func() string {
		if nxt >= len(s) { // reached the end.
			return nil
		}
		i := nxt
		nxt += 1
		return func() string {
			return s[i]
		}
	}
}

func main() {
	s := []string{"Hi", "Sabu", ",", "How", "are", "you", "?"}

	itr := makeItr(s)
	for c := itr(); c != nil; c = itr() {
		fmt.Println(c())
	}
}
