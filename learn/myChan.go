package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	ch := make(chan string, 3)
	go producer(ch)
	go customer(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("main end")

}

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func customer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
