package main

import "fmt"

func main() {

	var a int = 10;

	fmt.Printf("a address : %x", &a)
	fmt.Println("")	
	fmt.Printf("a val : %d", *&a)

}