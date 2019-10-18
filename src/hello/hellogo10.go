package main

import "fmt"

func f(x int, y int) (int, int) {
	return x+y, x-y
}

func main() {
	
	var x int = 10
	var y int = 11

	r1, r2 := f(x, y)

	fmt.Printf("x=%d, y=%d, x+y=%d, x-y=%d", x, y, r1, r2)
}