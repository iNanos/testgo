package main

import (
	"fmt"
	"strings"
)

func main() {
	var b bool = true
	var a = false
	c := 1

	d := 2

	str1 := "nanoxi\r\tong"
	str2 := "nano \t"
	str3 := "is \n"
	str4 := "\rxiongmengfei"

	fmt.Println("\nbefore replace" + str1)
	strings.Replace(str1, "xi", "", -1)
	fmt.Println("\nafter replace" + str1)

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 3))


	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(str1 + str2 + str3 + str4)

}