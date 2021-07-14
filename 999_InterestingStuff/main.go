// This code is used to reason about closures.
// Basically key here is when you use a closure you should remember that
// closures capture variables by reference. This example is from
// https://stackoverflow.com/questions/19957323/go-closure-variable-scope

package main 
import "fmt"


func makeIterator(sx []string) func() func() string {

	i := 0
	return func() func() string {
		if i == len(sx) {
			return nil
		}

		j := i
		i = i + 1
		return func() string {
			return sx[j]
		}
	}
}
func main() {

	i := makeIterator([]string{"Hello", "World", "This", "is", "Sabu"})

	for c := i(); c != nil; c = i() {
		fmt.Println(c())
	}
}