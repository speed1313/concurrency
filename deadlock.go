package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var mu2 sync.Mutex
	go func() {
		mu.Lock()
		time.Sleep(1 * time.Millisecond)
		mu2.Lock()
		fmt.Println("hello from p1")
		mu.Unlock()
		mu2.Unlock()
	}()
	go func() {

		mu2.Lock()
		time.Sleep(1 * time.Millisecond)
		mu.Lock()
		fmt.Println("hello from p2")
		mu2.Unlock()
		mu.Unlock()
	}()
	time.Sleep(5 * time.Second)
}
