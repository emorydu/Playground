# Gopher08

- 切片原理

1. 数组
```go
// Go数组是一个固定长度、容纳同构类型元素的连续序列，Go数组：类型+长度
// Go中传递数组是纯粹的值拷贝，需要考虑性能损耗
var a [8]int
var b [8]byte
var c [9[]int
```

2. 切片
```go
// 切片是数组的“描述符”
// 切片作为函数参数传递避免了较大性能损耗
// 切片的描述符是固定大小的、无论底层数组元素类型多大，切片打开窗口多长
// 新创建的切片与原切片同样是共享底层数组的，并且通过新切片对数组的修改也会反映到原切片中
// 切片作为参数传递带来的性能损耗都是很小且恒定的，甚至小到可以忽略不计

// src/runtime/slice.go
type slice struct {
	array unsafe.Pointer  // 指向下层数组某元素的指针，该元素是切片的起始元素
	len int // 切片的长度，切片中当前元素的个数
	cap int // 切片的最大容量 cap >= len
}

// 创建切片实例
s := make([]byte, 5)
```

3. 切片特性：动态扩容
```go
// append会根据切片的需要，在当前底层数组容量无法满足情况下，动态分配新的数组
// 新数组长度会按一定算法扩展。新数组建立后，append会吧旧数组中的数据复制到新
// 数组中，之后新数组便成为切片的底层数组，旧数组被GC
package main

import "fmt"

func main() {
	u := []int{11, 12, 13, 14, 15}
	fmt.Println("array:", u)                                     // [11, 12, 13, 14, 15]
	s := u[1:3]                                                  // [11, 12]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=2, cap=4): [12, 13]

	s = append(s, 24)
	fmt.Println("after append 24, array:", u)                                     // [11, 12, 13, 24, 15]
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=3, cap=4): [12, 13, 24]

	s = append(s, 25)
	fmt.Println("after append 25, array:", u)                                     // [11, 12, 13, 24, 25]
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=4, cap=4): [12, 13, 24, 25]

	s = append(s, 26)
	fmt.Println("after append 26, array:", u)                                     // [11, 12, 13, 24, 25]
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // slice(len=5, cap=8): [12, 13, 24, 25, 26]

	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u) // after reassign 1st elem of slice, array: [11 12 13 24 25]

	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // after reassign 1st elem of slice, slice(len=5, cap=8): [22 13 24 25 26]
}
```

4. 尽量使用cap参数创建切片
```go
// 如果可以预估出切片底层数组需要承载的元素数量，强烈建议在创建切片时带上cap参数
// 可以提升append的平均操作性能，减少或消除因动态扩容带来的性能损耗
```