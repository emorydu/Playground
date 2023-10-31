# Gopher01

- 追求简单，少即是多

少量的关键字

推崇“最小方式”思维方式

向后兼容

语言简单

- 偏好组合，正交解耦

无类型体系

1. 类型之间独立，没有子类型概念
2. 类型拥有自己的方法集合，类型定义与方法实现正交独立
3. 接口与其实现之间隐式关联
4. 包之间相互独立，没有子包的概念

垂直组合

类型嵌入，可以将已经实现的功能嵌入新类型中，快速满足新类型的功能需求
```go
// example 01.txt
type poolLocal struct {
	private interface{}
	shared []interface{}
	Mutex // type embedding
	pad [128]byte
}
// poolLocal拥有了Mutex类型的Lock以及Unlock方法

// example 02
// 聚合，用于组合成为大接口（方法集合）
type ReadWriter interface {
	Reader
	Writer
}
// 接口与实现者之间关系隐式，让程序各部分间耦合度降至最低，也是连接各部分的“纽带”
// 
// 在不经意间满足依赖抽象、里氏替换、接口隔离等设计原则
```

水平组合

通过interface将程序的各个部分组合在一起的方法
```go
func ReadAll(r io.Reader) ([]byte, error)

func Copy(dst Writer, src Reader) (written int64, err error)
```

- 原生并发，轻量高效

并发有关结构、并行有关执行

1. Go采用轻量级协程并发模型，Go应用面向多核硬件更具可扩展性
2. 提供支持并发的语法元素和机制
```go
// C/C++/Java
/**
执行单元：线程
调用库函数或调用对象方法进行创建/销毁
并发线程间通信：基于操作系统提供的IPC机制（共享内存、Socket、Pipe、并发保护的全局变量...）
 */

// Go
/**
执行单元：goroutine
go+函数调用；函数退出goroutine退出
并发goroutine通信：通过语言内置channel传递消息或实现同步，通过select实现多路channel的并发控制
 */
```
3. 并发有关结构，是一种将一个程序分解为多个小片段并且每个小片段都可以独立执行的程序设计方法；并发程序的小片段之间一般存在通信联系并且通过通信相互协作。
4. 并行是有关执行的，表示同时进行一些计算任务。
   
[Rob Pike](https://go.dev/talks/2012/waza.slide)

5. 并发程序的结构设计不要局限于在单核情况下处理能力的高低，而要在多核情况下充分提升多核利用率、获得性能的自然提升为最终目的。

Go工具链
1. go build / go run
2. go list / go get / go mod xx
3. go fmt / gofmt
4. go doc / godoc
5. go test
6. go vet
7. go tool pprof / go tool trace
8. go tool fix
9. go bug


