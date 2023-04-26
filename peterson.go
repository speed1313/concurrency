package main

import (
	"fmt"
	"time"
)

var turn int
var flag [2]bool

func main() {

	go peterson(0)
	go peterson(1)
	time.Sleep(1 * time.Millisecond)
}

func peterson(i int) {
	for {
		fmt.Println("i:", i)
		flag[i] = true
		turn = 1 - i
		for flag[1-i] && turn == 1-i {
			//do nothing
		}
		time.Sleep(1000 * time.Nanosecond)
		//critical section
		flag[i] = false

	}
}
