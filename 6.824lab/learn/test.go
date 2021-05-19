package main

import "fmt"

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simpleTwo() func(a, b int) int {
	f := func(a, b int) int {
		return a - b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	instanceTwo := simpleTwo()
	fmt.Println(instanceTwo(1000, 9))

	fmt.Println("*************************")
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Goer"))
	fmt.Println(b("!"))
}
