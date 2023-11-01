package main

import (
	"fmt"
	"time"
)

func Exception1() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}

func Success1() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		i := i // here
		v := v // here
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}

func Success2() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 10)
}
