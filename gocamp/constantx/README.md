# Gopher06

- 使用无类型常量
```go
// src/io/io.go
const (
	SeekStart = 0
	SeekCurrent = 1
	SeekEnd = 2
)
```

1. Go中，两个类型即使拥有相同的底层类型，也仍然是不同的数据类型
```go
type MyInt int

func main() {
  var a int = 5
  var b MyInt = 5
  // fmt.Println(a == b) // compiler error
  fmt.Println(a == int(b))
}
```
```go
type myInt int
const n myInt = 13
const m int = n + 15  // compiler error

func main() {
	var a int = 5
	fmt.Println(a + n)  // compiler error
}
```
```go
const (
  a = 5
  pi = 3.14159
  s = "Hello, Emory.Du:)"
  c = 'c'
  b = true
)

type myInt int
type myFloat float32
type myString string

func main() {
  var j myInt = a
  var f myFloat = pi
  var str myString = s
  var e float64 = a + pi
  
  fmt.Println(j)
  fmt.Println(f)
  fmt.Println(str)
  fmt.Printf("%T, %v\n", e, e)
}
```
2. iota定义类型安全/无类型枚举