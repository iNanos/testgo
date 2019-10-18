package main

import "fmt"

func main() {

	var numbers = make([]int, 3, 5)

	fmt.Printf("len=%d, cap=%d slice=%v", len(numbers), cap(numbers), numbers)

}