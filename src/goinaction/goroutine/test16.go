package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines
	counter3 int
	// wg is used to wait for the program to finish
	wg3 sync.WaitGroup
	// mutes is used to define a critical section of code
	mutex sync.Mutex
)

func main() {
	// add a count of two, one for each goroutine
	wg3.Add(2)

	// create two goroutine
	go incCounter1(1)
	go incCounter1(2)

	// wait for the goroutines to finish
	wg3.Wait()
	fmt.Printf("Final Counter: %d\\n", counter3)
}

// increments the package level Counter variable
// using the mutex to synchronize and provide safe access
func incCounter1(id int) {

	defer wg3.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入临界区
		mutex.Lock()
		{
			// 捕获counter的值
			value := counter3
			// 当前goroutine从线程退出，并放回到队列
			runtime.Gosched()
			//
			value++
			// 将值保存到counter
			counter3 = value
		}
		// 释放锁，允许其它正在等待的goroutine进入临界区
		mutex.Unlock()
	}
}
