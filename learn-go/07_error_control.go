package main

import "fmt"

// Defer
// Defers running a function until the surrounding function returns. The arguments are evaluated immediately, but the function call is not ran until later.
func main() {
	defer fmt.Println("done")
	fmt.Println("working")
}

// Deferring functions
// Lambdas are better suited for defer blocks.
func main() {
	defer func() {
		fmt.Println("done")
	}()
	fmt.Println("working")
}
