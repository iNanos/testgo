package main

import (
	"fmt"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err %s", err)
		}
	}()

	defer func() {
		fmt.Println("three")
		panic("three")
	}()

	defer func() {
		fmt.Println("two")
		panic("two")
	}()

	fmt.Println("one")
	panic("one")

}
