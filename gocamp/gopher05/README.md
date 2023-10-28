# Gopher05

- 声明、类型、语句与控制结构

1. 使用一致的变量声明形式
2. 在变量声明形式的选择上应尽量保持项目内一致
3. 包级变量：package级别可见，如果包级变量是可导出的，则视为全局变量
4. 局部变量：函数或者方法体内声明的变量、返回值、函数参数，仅函数/方法内可见
```go
// 变量使用之前必须先声明
var a int32
var s string = "string"
var i = 13
n := 17

var (
	crlf = []byte("\r\n")
	colonSpace = []byte(": ")
)
```

5. 声明并同时显式初始化
```go
// src/io/pipe.go
var ErrClosedPipe = errors.New("io: read/write on closed pipe")

// src/io/io.go
var EOF = errors.New("EOF")
var ErrShortWrite = errors.New("short write")

// 具有默认类型
var a = 17  // int
var f = 3.14  // float64



// 推荐做法
var a = int32(17)
var f = float32(3.14)

var (
    a = int32(17)
    f = float32(3.14)
)
```

6. 声明但延迟初始化
```go
var number int

...

number = 24
```
7. 声明类聚与就近原则

```go
// src/net/http/server.go

var (
	bufioReaderPool   sync.Pool
	bufioWriter2kPool sync.Pool
	bufioWriter4kPool sync.Pool
)

var copyBufPool = sync.Pool {
	New: func() interface{} {
		b := make([]byte, 32*1024)
		return &b
	},
}

// src/net/net.go
var (
	aLongtimeAgo = time.Unix(1, 0)
	noDeadline   = time.Time{}
	noCancel     = (chan struct{})(nil)
)

var threadLimit chan struct{}

// src/net/http.request.go
var ErrNoCookie = errors.New("http: named cookie not present")

func (r *Request) Cookie(name string) (*Cookie, error) {
	for _, c := range readCookies(r.Header, name) {
		return c, nil
	}
		
	return nil, ErrNoCookie
}
```

- 局部变量

1. 延迟初始化局部变量
```go
// src/strings/replace.go
func (r *byteReplacer) Replace(s string) string {
	var buf []byte
	for i := 0; i < len(s); i++ {
		b := s[i]
		if r[b] != b {
		  if buf == nil {
			buf = []byte(s)
		  }
		  buf[i] = r[b]
		}
	}
	if buf == nil {
          return s
	}
	return string(buf)	
}
```
2. 声明且显示初始化局部变量，建议短变量声明
```go
a := 17
f := 32.4
s := "Hello, Emory:)"

f := float32(1.32)
a := int32(15)
s := []byte("hello, emory:)")
```

- 函数/方法单一职责
```go
func (r *Resolver) resolveAddrList(ctx context.Context, op, network,
                            addr string, hint Addr) (addrList, error) {
    ...
    var (
        tcp      *TCPAddr
        udp      *UDPAddr
        ip       *IPAddr
        wildcard bool
    )
    ...
}
```


