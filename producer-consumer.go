package main

import (
	"fmt"

	"time"
)

const BUFFER_SIZE = 10

func main() {
	var buffer [BUFFER_SIZE]int
	var counter int
	var in int
	var out int

	go func() {
		for {

			for counter == BUFFER_SIZE {
				//do nothing
			}
			buffer[in] = in
			in = (in + 1) % BUFFER_SIZE
			counter++
			fmt.Println("counter:", counter)
		}
	}()
	go func() {
		for {
			for counter == 0 {
				//do nothing
			}
			nextConsumed := buffer[out]
			//fmt.Println("nextConsumed:", nextConsumed)
			out = (out + 1) % BUFFER_SIZE
			counter--

			nextConsumed++
		}
	}()

	time.Sleep(1000000 * time.Nanosecond)
}
