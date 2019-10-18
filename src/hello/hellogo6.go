package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0;
	for i, val := range s {
		sum += val
		fmt.Printf("index=%d, value=%d\n", i, val)
	}

	ch <- sum
}

func main() {

	slice := []int {1,2,3,4,5,6,7,8,9}

	c := make(chan int, 100)

	go sum(slice[:len(slice)/2], c)

	go sum(slice[len(slice)/2:], c)

	x, y := <-c, <-c

	fmt.Println("\n\n")
	fmt.Printf("x(%d) + y(%d) = %d", x, y, x + y)

	close(c)

	dc := make(chan int, 2)
    dc <- 1
    dc <- 2
    // dc <- 3s
    fmt.Println("\n\n")
    fmt.Println(<-dc)
	fmt.Println(<-dc)

	close(dc)
}