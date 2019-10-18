package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}


//名为DivideError的struct实现了 error接口
//go 的 接口 实现是隐式的 ， 不需要implements
//go 的 实现了接口的所有方法, 即为实现了该接口
func (de DivideError) Error() string {
	strFormat := "the divider can't be 0. dividee:%d. divider:0"
	return fmt.Sprintf(strFormat, de.dividee)
}

func Devide(dividee int, divider int) (result float32, errStr string) {
	if divider == 0 {
		err := DivideError{
			dividee:dividee,
			divider:divider,
		}
		errStr = err.Error()
		return 
	} else {
		result := float32(dividee) / float32(divider)
		return result, ""
	}
}

func main() {

	if result, err := Devide(100, 3); err == "" {
		fmt.Println("100/3=", result)
	} else {
		fmt.Println("100/3=", err)
	}

	if result, err := Devide(100, 0); err != "" {
		fmt.Println("100/0=", err)
	} else {
		fmt.Println("100/0=", result)
	}

}