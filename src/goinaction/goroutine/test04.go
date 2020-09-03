package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 用来等待程序完成
var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Waiting to Finish")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("Terminating program")

}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)

}
