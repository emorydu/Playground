# Gopher12

- 初始化依赖规则

包级别变量声明语句的表达式求值顺序是由初始化依赖规则决定的
1. Go中，包级别变量的初始化按照变量声明的先后顺序进行
2. 如果某个变量的初始化表达式中直接或间接依赖其他变量，那么变量1的初始化顺序排在变量2的后面
3. 未初始化的且不含有对应初始化或初始化表达式不依赖任何未初始化变量的变量，称为`ready for initialization`
4. 包级别变量的初始化是逐步进行的，每一步就是按照变量声明顺序找到下一个`ready for initialization`
5. 位于同一包内但不同文件中的变量的声明顺序依赖编译器处理文件的顺序：先处理的文件中的变量的声明顺序先于后处理的文件中的变量
```go
// 1
var (
  a = c + b
  b = f()
  c = f()
  d = 3
)

func f() int {
  d++
  return d
}

func main() {
  fmt.Println(a, b, c, d)
}


// 2
var (
  a = c + b
  b = f()
  _ = f()
  c = f()
  d = 3
)

func f() int {
  d++
  return d
}

func main() {
  fmt.Println(a, b, c, d)
}


// 3
var (
  a    = c
  b, c = f()
  d    = 3
)

func f() (int, int) {
  d++
  return d, d + 1
}

func main() {
  fmt.Println(a, b, c, d)
}
```

- 普通求值顺序

普通求值顺序用于规定表达式操作数中的函数、方法以及channel操作的求值顺序

Go规定表达式操作数中的所有函数、方法以及channel操作按照`从左到右`的次序进行求值

- 赋值语句

`从左到右求值`  / `从左到右赋值`

- switch/select表达式求值
```go
// switch
func Expr(n int) int {
  fmt.Println(n)
  return n
}

func main() {
  switch Expr(2) {
  case Expr(1), Expr(2), Expr(3):
    fmt.Prinltn("enter into case1")
    fallthrough
  case Expr(4):
    fmt.Println("enter into case2")
  }
}
```

```go
// select
package main

import (
	"fmt"
	"time"
)

func getAReadOnlyChannel() <-chan int {
	fmt.Println("invoke getAReadOnlyChannel")
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getASlice() *[5]int {
	fmt.Println("invoke getASlice")
	var a [5]int
	return &a
}

func getAWriteOnlyChannel() chan<- int {
	fmt.Println("invoke getAWriteOnlyChannel")
	return make(chan int)
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

func main() {
	select {
	case (getASlice())[0] = <-getAReadOnlyChannel():
		fmt.Println("recv something from a readonly channel")
	case getAWriteOnlyChannel() <- getANumToChannel():
		fmt.Println("send something to a write-only channel")
	}
}
// select执行开始时，首先所有的cases表达式都会被按出现的先后顺序求值一遍
// 位于cases等号左边的从channel接收数据的表达式RecvStmt不会被求值
```

- 总结
1. 包级别变量声明语句中的表达式求值顺序由变量的声明顺序和初始化依赖关系决定，并且包级变量表达式求值顺序优先级最高
2. 表达式操作数中的函数、方法及channel操作按普通求值顺序，即从左到右的次序进行求值
3. 赋值语句求值分为两个阶段：先按照普通求值规则对等号左边的下标表达式、指针解引用表达式和等号右边的表达式中的操作数进行求值，然后按从左到右的顺序对变量进行赋值