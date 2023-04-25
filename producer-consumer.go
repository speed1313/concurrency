package main

import (
	"fmt"
	"math/rand"
	"time"
)

const BUFFER_SIZE = 10

func main(){
	var buffer [BUFFER_SIZE]int
	var counter int
	var in int
	var out int

	go func(){
		for{
			nextProduces := rand.Int()
			for(counter == BUFFER_SIZE){
				//do nothing
			}
			buffer[in] = nextProduces
			in = (in + 1) % BUFFER_SIZE
			counter++
		}
	}()
	go func(){
		for{
			for(counter == 0){
				//do nothing
			}
			nextConsumed := buffer[out]
			out = (out + 1) % BUFFER_SIZE
			counter--
			nextConsumed ++
		}
	}()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("counter:", counter)
}