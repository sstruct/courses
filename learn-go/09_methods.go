package main

import "math"

// Receives
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64  {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

v: = Vertex{1, 2}
v.Abs()

// Mutation
func (v * Vertex) Scale(f float64)  {
	v.X = v.X * f
	v.y = v.Y * f
}

v := Vertex{6, 12}
v.Scale(0.5)