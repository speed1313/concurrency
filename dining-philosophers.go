package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	// atomic chopsticks
	chopsticks := make([]sync.Mutex, 5)

	for i := 0; i < 5; i++{
		i := i
		go func(i int){
			for{
				// deadlock can happen more easily by increasing the number of Locks
				chopsticks[i].Lock()
				chopsticks[(i+1)%5].Lock()
				//chopsticks[(i+2)%5].Lock()
				//chopsticks[(i+3)%5].Lock()
				// eating
				fmt.Println("eating: ", i)
				time.Sleep(1 * time.Millisecond)
				chopsticks[i].Unlock()
				chopsticks[(i+1)%5].Unlock()
				//chopsticks[(i+2)%5].Unlock()
				//chopsticks[(i+3)%5].Unlock()
				// thinking
			}
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
}