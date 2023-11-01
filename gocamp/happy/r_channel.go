package main

import (
	"fmt"
	"time"
)

func Channel() {
	// channel在Go运行时表示为一个channel的描述符指针，因此channel的指针副本也指向原channel

	// 当channel作为range表达式类型时，for range最终以阻塞读的方式阻塞在channel表达式上，
	// 即便是带缓冲的channel亦是如此：当channel中无数据时，for range也会阻塞在channel上，直到channel关闭
	var c = make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func NilChannel() {
	var c chan int

	// for range将永远阻塞在这个nil channel上，直到runtime发现程序陷入deadlock状态，panic
	for v := range c {
		fmt.Println(v)
	}
}
func main() {
	NilChannel()
}
