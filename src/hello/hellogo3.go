package main

import "fmt"

func main() {

	var input int

	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}