package main

import "fmt"

type Printer func(args ...interface{}) (int, error)

func PrintToStd(args ...interface{}) (int, error) {
	return fmt.Printf("%v\n", args)
}

func main() {
	var p Printer

	p = PrintToStd
	p("Hello, ", "Nice to meet you!")

	fmt.Printf("%T\n", PrintToStd)
}
