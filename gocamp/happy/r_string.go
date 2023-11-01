package main

import "fmt"

func String() {
	// string在Go运行时内部表示：struct{*byte, len}, 且string本身是immutable
	// 与slice类似
	var s = "中文"
	for i, v := range s { // 每次循环的单位是一个rune，不是一个byte
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
}
