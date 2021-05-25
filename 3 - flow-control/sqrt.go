package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	lastZ := 0.0
	for {
		z -= (z*z - x) / (2*z)
		if (z-lastZ) < 0.01 {
			return z
		}
		lastZ = z
		fmt.Println(z)
	}
	return z
	
}

func main() {
	fmt.Println(Sqrt(2))
}
