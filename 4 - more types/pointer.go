package main

import "fmt"

/*
The type *T is a pointer to a T value. Its zero value is nil.

var p *int

The & operator generates a pointer to its operand.

i := 42
p = &i

Unlike C, Go has no pointer arithmetic.
*/

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
