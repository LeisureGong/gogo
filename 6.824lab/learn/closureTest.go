package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 高阶函数： 将函数类型分别作为函数的参数和返回值
type calculateFn func(x, y int) (int, error)

// genCalculator 定义一个匿名的、calculateFn 类型的函数并把它作为
// 结果值返回，而这个匿名的函数就是一个闭包函数。
//
// 匿名函数里面使用的变量 op  既不代表它的任何参数或结果也不是它自己// 声明的，而是定义它的 genCalculator 函数的参数，所以是一个自由变// 量。
//
// 这个自由变量 op 究竟代表了什么，这一点并不是在定义这个闭包函数的
// 时候确定的(其实，这时候知道该变量的类型)，而是在 genCalculator
// 函数被调用的时候确定的。
func getCalculator(op operate) calculateFn {
	return func(x, y int) (int, error) {
		// op是从外部获取的
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func functionalB() {
	x, y := 98, 23

	op := func(x, y int) int {
		return x - y
	}
	sub := getCalculator(op)

	ret, err := sub(x, y)
	if err != nil {
		fmt.Printf("计算错误：%v\n", err)
		return
	}
	fmt.Printf("result: %v\n", ret)
}

func main() {
	functionalB()
}
