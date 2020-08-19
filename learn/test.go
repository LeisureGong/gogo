package main

import "fmt"

func main() {
	person := []string{"Tom", "Aaron", "Jhon"}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(person), cap(person), person)
	fmt.Println("")

	// 循环
	for k, v := range person {
		fmt.Printf("person[%d] : %s\n", k, v)
	}

	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	for _, name := range person {
		fmt.Println("name: ", name)
	}
}
