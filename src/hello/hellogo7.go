package main

import "fmt"

func main() {

	var count int = 1
	var flag bool
	for count < 100 {
		count++
		flag = true
		for tmp:=2;tmp<count;tmp++ {
			if count%tmp==0 {
				flag = false
				break
			}
		}

		if flag==true {
			fmt.Printf("\ncount=%d sushu", count)
		} 
		//else {
		//	fmt.Printf("\ncount=%d is not sushu", count)
		//}
	}

fmt.Println("")

	for i:=1;i<10;i++ {
		for j:=1;j<10;j++ {
			fmt.Printf("%2d *%2d = %2d ", i, j, i*j)
		}
		
		fmt.Println("")

	}

}