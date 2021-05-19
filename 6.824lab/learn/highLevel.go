package main

import (
	"errors"
	"fmt"
)

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}

	return op(x, y), nil
}

func functional() {

	var op operate
	//fmt.Println("op = ", op)
	//op(1, 2)

	op = func(x, y int) int { return x + y }
	fmt.Println("op = ", op)

	ret, err := calculate(2, 3, op)
	if err != nil {
		fmt.Printf("calculate failed: %v\n", err)
	}
	fmt.Printf("result: %v\n", ret)

}

func main() {
	functional()
}
