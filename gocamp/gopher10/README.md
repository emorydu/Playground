# Gopher10

- 字符串

1. 字符串字面量
```go
package main

import "fmt"

const (
	s = "string constant"
)

func main() {

	var (
		sv = "string variable"
	)

	fmt.Printf("%T\n", s)                          // string
	fmt.Printf("%T\n", sv)                         // string
	fmt.Printf("%T\n", "temporary string literal") // string
}
```
2. string类型的数据是不变的

Go编译器会为切片变量重新分配底层存储而不是共用string的底层存储，因此对切片的修改并未对原string的数据产生任何影响
```go
package main

import "fmt"

func main() {
	var name = "Emory.Du"
	fmt.Println("original string:", name)

	nl := []byte(name)
	nl[0] = 'e'
	fmt.Println("slice:", string(nl)) // emory.Du
	fmt.Println("after reslice, the original string is:", name)
}
```
3. 零值可用
```go
var s string
fmt.Println(s)  // ""
fmt.Println(len(s)) // 0

// 获取长度的时间复杂度是O(1)
/*
Go string类型是不变的（数据不会变、长度也不会变）所有的表示都是在string的内部结构体中作为值
 */

// 支持+、+=操作
s := "Emory"
s = s + ".Du"
s += " :)"

// 支持各种比较关系操作符：==、!=、>=、 <=、 >、 <
/*
Go string是不可变的，因此如果两个字符串的长度不相同，
那么无须比较具体字符串数据即可断定两个字符串是不同的。
如果长度相同，则要进一步判断数据指针是否指向同一块底层存储数据。
如果相同，则两个字符串是等价的；如果不同，则还需进一步比对实际的数据内容
 */
```

- 字符串内部表示
```go
// src/runtime/string.go
type stringStruct struct {
  str unsafe.Pointer
  len int
}

func rawstring(size int) (s string, b []byte) {
  p := mallocgc(uintptr(size), nil, false)
  stringStructOf(&s).str = p
  stringStructOf(&s).len = size
  *(*slice)(unsafe.Pointer(&b)) = slice{p, size, size}
  return
}
```
[图片描述](./imgs/01.png) / [文字解释](./txt/01.txt)

- 高效使用string
1. fmt.Sprintf ⭐️⭐️：适用于构建特定格式
2. strings.Join ⭐️⭐️⭐️：输入的字符串是[]string承载的
3. strings.Builder ⭐️⭐️⭐️⭐️⭐️（预初始化） / ⭐️⭐️⭐️（未预初始化）
4. bytes.Buffer ⭐️⭐️⭐️

- 字符串高效转换

string到[]rune以及string到[]byte的转换，这两个转换也是可逆的，也就是说string和[]rune、[]byte可以双向转换

无论是string转slice还是slice转string，转换都是要付出代价的，这些代价的根源在于string是不可变的，运行时要为转换后的类型分配新内存