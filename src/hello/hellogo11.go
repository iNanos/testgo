package main

import "fmt"

func max(param1 int, param2 int) int {
	if param1 < param2 {
		return param2
	} else if param1 == param2 {
		return param1
	} else {
		return param2
	}
}

func main() {

	var param1 int = 10

	var param2 int = 12

	fmt.Printf("max=%d", max(param1, param2))

}