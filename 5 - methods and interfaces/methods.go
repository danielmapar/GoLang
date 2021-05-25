package main

import (
	"fmt"
	"math"
)

/*
A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.
*/

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}