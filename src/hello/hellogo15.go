package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os len=", len(os.Args))
	for k, v := range os.Args {
		fmt.Printf("k=%s -> v=%s", k, v)
	}

	fmt.Println("")
	var arr = [3]int{4,5,6}
	for i, v := range arr {
		fmt.Printf("index=%d, val=%d\n", i, v)
	}

	fmt.Println("")
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["nano1"] = "nano1"
	map1["nano2"] = "nano2"
	map1["nano3"] = "nano3"
	map1["nano4"] = "nano4"
	for k := range map1 {
		fmt.Printf("k=%d, v=%d", k, map1[k])
		fmt.Println("")
	}
}