package main

import (
	"fmt"
)

type Human interface {
	name() string
	age() int
}

type Women struct {
}

func (women Women) name() string {
	return "hanmeimei"
}

func (women Women) age() int {
	return 20
}


type Men struct {
}

func (men Men) name() string {
	return "lilei"
}

func (men Men) age() int {
	return 22
}
	
func main() {
	var human1 Human = new(Women)
	fmt.Println(human1.name())
	human1 = new(Men)
	fmt.Println(human1.name())
}