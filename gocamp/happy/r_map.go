package main

import "fmt"

func Map() {
	// map类型作为range表达式时，会得到一个map的内部表示的副本
	// for-range 对map副本的操作即对源map的操作
	var m = map[string]int{
		"Emory.Du": 24,
		"Smith":    25,
		"Josh":     23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "Josh")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)

	// tips:
	// for range无法保证每次迭代的元素次序是一致的。
	// 同时，如果在循环的过程中对map进行修改，那么这样修改的结果是否会影响后续迭代过程也是不确定的
}
