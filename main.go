package main

import (
	"sync"
)

func producer(ch chan int, i int){
	ch <- i * 2
}


func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer 
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
}
