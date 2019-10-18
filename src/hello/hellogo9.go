package main

import "fmt"

func avg(arr [5]int) float32 {
	var sum int;
	for i:=0;i<5;i++ {
		sum += arr[i]
	}

	return float32(sum) / float32(len(arr))	
}

func main() {
	var arr = [5]int{3,19,28,100,6}
	fmt.Printf("avg=%f", avg(arr))

	fmt.Println("")
	var arr1 = []float32{1.9, 2.0, 4.1}
	printlnArrLen(arr1)


}

func printlnArrLen(param []float32) {
	fmt.Printf("the array length = %d", len(param))
	// fmt.Println("")
	// fmt.Printf("the array length = %d", len)
}