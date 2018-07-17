package main

import "fmt"

// strings
str := "hello"
str := `Multiline
string`

// Numbers
num := 3
num := 3.
num := 3 + 4i
num := byte('a')

// Arrays have a fixed size.
numbers := [...]int{0, 0, 0, 0, 0}

// Pointers
func main()  {
	b := *getPointer()
	fmt.Println("value is ", b)
}

func getPointer() (myPointer *int) {
	a := 234
	return &a
}

// Type conversions
i := 2
f := float64(i)
u := unit(i)

// Slices have a dynamic size, unlike arrays.
slice := []int{2, 3, 4}
slice := [byte]("hello")

