package main

import (
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s Student) GetName() string {
	fmt.Println(s.Name)
	return s.Name
}

func main() {
	var stu = new(Student)
	stu.Age = 22
	stu.Name = "gonglei"
	stu.GetName()

}
