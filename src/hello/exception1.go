package main

import "fmt"

func main() {

	fmt.Println("外层开始")

	defer func() {

		fmt.Println("外层defer recover")
		if err := recover(); err != nil {
			fmt.Println("外层recover到异常, err: %s", err)
		} else {
			fmt.Println("外层没recover到异常")
		}

		fmt.Println("外层 recover完成")
	}()

	fmt.Println("f1即将异常")

	f1()

	fmt.Println("f1异常后")

	defer func() {
		fmt.Println("f1异常后 defer")
	}()
}

func f1() {
	fmt.Printf("内层开始")
	defer func() {
		fmt.Println("内层recover前的defer")
	}()

	defer func() {
		fmt.Println("内层准备recover")
		if err := recover(); err != nil {
			fmt.Println("recover get err : %s", err)
		}

		fmt.Println("内层recover完毕")
	}()

	defer func() {
		fmt.Println("内层异常前 recover后的 defer")
	}()

	panic("内层发生异常")

	defer func() {
		fmt.Println("内层异常后的 defer")
	}()

	fmt.Println("内层异常后的语句")
}
