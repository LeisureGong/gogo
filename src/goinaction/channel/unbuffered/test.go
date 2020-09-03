package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func player(name string, court chan int) {
	// 通知main函数
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}

func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)

	// start
	go player("gonglei", court)
	go player("chengruimin", court)

	court <- 1
	wg.Wait()
}
