package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg2      sync.WaitGroup
)

func main() {

	wg2.Add(2)

	go doWork("A")
	go doWork("B")

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wg2.Wait()

}

func doWork(name string) {
	defer wg2.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutdting %s Down\n", name)
			break
		}
	}
}
