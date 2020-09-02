package main

import "fmt"

type user struct {
	name  string
	email string
}

// notifier是一个定义了通知类型为的接口
type notifier interface {
	notify()
}

// notify是使用指针接收者实现的方法
func (u user) notify() {
	fmt.Printf("发送邮件给%s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"宫雷", "leisure_gong@163.com"}
	sentNotification(u)
}

// sentNotification接受一个实现notifier接口的值，并发送通知
func sentNotification(n notifier) {
	n.notify()
}
