package main

import (
	"fmt"
	"time"
	"runtime"
	"bytes"
	"strconv"
)

func getGoRoutineId() {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err==nil {
		fmt.Println("goroutine_id = " + strconv.FormatUint(n, 16))
	} else {
		fmt.Println(err)
	}
}

func say(s string) {
	for i:= 0; i<5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("hello, " + s)
	}

	getGoRoutineId()
	
}

func main() {
	go say("goroutine, nanoxiong")
	say("nanoxiong")
}