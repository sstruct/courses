package main

import "fmt"

// Defining
type Vertex struct {
	X int
	Y int
}

func main()  {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
}

// Literals
v := Vertex{X: 1, Y: 2}
// Field names can be omitted
v := Vertex{1, 2}
// Y is implicit
v := Vertex{X: 1}

// Pointers to structs
// Doing v.X is the same as doing (*v).X, when v is a pointer.
v := &Vertex{1, 2}
v.X = 2
