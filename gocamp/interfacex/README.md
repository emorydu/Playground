# Gopher04

- 命令简单一致
1. 包名：小写单个单词命名 / 尽量与包导入路径的最有一个路径分段保持一致
2. 在包中兼顾该包导出的标识符（变量、常量、类型、函数...）命名，不再包含包名
```go
strings.Reader
strings.Writer

bytes.Buffer
bytes.NewBuffer
...
```

3. 包级变量：
4. 局部变量：函数或方法内的变量、函数或方法的参数、返回值
5. LowerCamelCase
```go
package main

import (
	"fmt"
)
// 循环和条件变量多采用单个字母命名
// 函数/方法的参数和返回值变量以单个单词或单个字母为主
// 由于方法在调用时会绑定类型信息，因此方法的命名以单个单词为主
// 函数多以多单词的复合词进行命名
// 类型多以多单词的复合词进行命名

type User struct{}

func main() {
	users := []*User{{}, {}}
    fmt.Println(users)
}

// 示例代码
for i, v := range s { ... } // slice

for k, v := range m { ... } // map

for v := range r { ... }  // channel

t := time.Now() // time
t := &Timer{} // timer
if t := md.typeMap[off]; t != nil { ... } // type

b := make([]byte, n)  // byte slice
b := new(bytes.Buffer)  // byte buffer
```

6. 常量：无需显示赋予类型，会自动推断，常量名不要包含类型信息
```go
// src/net/http/request.go
const (
	defaultMaxMemory = 32 << 20 // 32 Mb
)

const (
	deleteHostHeader = true
	keepHostHeader = false
)


// src/math.sign.go
const (
  PI4A = 7.85398125648498535156E-1  // 0x3fe921fb40000000,
  PI4B = 3.77489470793079817668E-8  // 0x3e64442d00000000,
  PI4C = 2.69515142907905952645E-15 // 0x3ce8469898cc5170,
)

// src/syscall/zerrors_linux_amd64.go

// 错误码
const (
    E2BIG           = Errno(0x7)
    EACCES          = Errno(0xd)
    EADDRINUSE      = Errno(0x62)
    EADDRNOTAVAIL   = Errno(0x63)
    EADV            = Errno(0x44)
    ...
)

// 信号
const (
    SIGABRT   = Signal(0x6)
    SIGALRM   = Signal(0xe)
    SIGBUS    = Signal(0x7)
    SIGCHLD   = Signal(0x11)
    ...
)
```
7. 接口：优先以单个单词命名，拥有唯一方法或通过多个唯一方法的接口组合而成的接口，使用方法名+er
8. 尽量构建小接口，使用接口组合构建程序
```go
// src/io/io.go

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriteCloser interface {
	Writer
    Reader
	Coser
}

```

9. 利用上下文，携带足够多的信息
```go
// bad
for index := 0; index < len(s); index++ {
	value := s[index]
    ...
}

// good
for i := 0; i < len(s); i++ {
	v := s[i]
	...
}
```