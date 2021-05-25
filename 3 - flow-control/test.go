package main

import (
	"fmt"
	"math"
	"strconv"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	sum := 1
	for i := 1; i < 1000; i++ {
		sum += i
	}
	// while loop
	for sum < 500000 {
		sum += 1
	}
	// for ever loop
	//for {
	//}


	result := ""
	if sum == 500000 {
		result = sqrt(float64(sum))
	}

	// Like for, the if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if
	lim := 5.0
	if v := math.Pow(float64(sum), 2); v < lim {
		result = strconv.FormatFloat(v, 'E', -1, 64)
	}

	//Variables declared inside an if short statement are also available inside any of the else blocks.

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	
	//sum := 1
	//for ; sum < 1000; {
	//	sum += sum
	//}
	fmt.Println(result)
}
