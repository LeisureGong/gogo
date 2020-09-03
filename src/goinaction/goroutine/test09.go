package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg1     sync.WaitGroup
)

func main() {

	wg1.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	wg1.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg1.Done()

	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}
