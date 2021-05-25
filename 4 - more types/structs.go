package main

import "fmt"

type Vertex struct {
	X int
	Y int
}


func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	/*
	To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
	However, that notation is cumbersome, so the language permits us instead to write just p.X, 
	without the explicit dereference.
	*/

	v2 := Vertex{1, 2}
	p := &v2
	p.X = 1e9
	fmt.Println(v2)

	var (
		v3 = Vertex{1, 2}  // has type Vertex
		v4 = Vertex{X: 1}  // Y:0 is implicit
		v5 = Vertex{}      // X:0 and Y:0
		p2  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v3, v4, v5, *p2, p2)
}
