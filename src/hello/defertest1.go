package main

import "fmt"

func f3() (r int) {
	defer func() {
		r++
	}()

	return 0
}

func main() {
	fmt.Printf("result %d", f3())
}
