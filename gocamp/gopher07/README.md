# Gopher07

- 尽量定义零值可用类型
```go
当通过声明或调用new为变量分配存储空间，或者通过复合文字字面量或调用make创建新值，
且不显式初始化时，Go为变量或值提供默认值。

零值：
整型：0
浮点类型：0.0
布尔类型：false
字符串类型：""
指针、interface、slice、channel、map、function：nil
Go零值递归：数组、结构体等类型的零值初始化就是对其组成元素逐一零值初始化
```

- 零值可用
1. slice append
```go
var zeros []int // nil
zeros = append(zeros, 1)
zeros = append(zeros, 2)
zeros = append(zeros, 3)
fmt.Println(zeros)  // not nil
```
2. nil pointer

```go
var p *net.TCPAddr
fmt.Println(p)  // <nil>
```

3. sync.Mutex
```go
var mu sync.Mutex
mu.Lock()
mu.Unlock()
```
4. bytes.Buffer
```go
type Buffer struct {
	buf []byte
	off int
	lastRead readOp
}
```
```go
// Go并非所有类型都是零值可用，且零值可用也有一定的限制
var s []int
s[0] = 12 // error
s = append(s, 12)

// map这样的原生类型没有提供对零值可用的支持
var m map[string]int
m["go"] = 1 // error

m := make(map[string][int])
m["go"] = 1

// 零值可用的类型要注意尽量避免值复制
var mu sync.Mutext
mcp := mu // error
foo(mu) // error

// 可以通过指针方式传递类型Mutex这样的类型
var mu sync.Mutex
foo(&mu)

// 保持与Go一致，给自定义类型一个合理的零值，并尽量保持自定义类型零值可用
```