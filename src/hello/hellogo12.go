package main

import (
	"fmt"
	"math"
)


func add(x1 int, x2 int) func(x3 int, x4 int)(int, int, int) {
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1+x2, x3+x4
	}
}


func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println("sqrt(16)=", getSquareRoot(16))

	add_func := add(1, 2)

	a, b, c := add_func(5, 6)
	d, e, f := add_func(3, 7)

	fmt.Printf("i=%d, 1+2=%d, 5+6=%d", a, b, c)
	fmt.Println("")
	fmt.Printf("i=%d, 1+2=%d, 3+7=%d", d, e, f)
}