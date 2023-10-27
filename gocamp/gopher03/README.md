# Gopher03

- 公认广泛的项目结构
1. 项目结构决定了项目内部包的布局及包依赖关系是否合理。
2. 影响外部项目对该项目中包的依赖与引用。

go1.21.3
```go
├── CONTRIBUTING.md
├── LICENSE
├── PATENTS
├── README.md
├── SECURITY.md
├── VERSION
├── api/
├── bin/
├── codereview.cfg
├── doc/
├── go.env
├── lib/
├── misc/
├── pkg/
├── src/
├── test/


// tips: 额外解释
...
├── internal/ // 用于存放无法被外部导入、仅Go项目自用的包
└── vendor/ // 存放Go项目自身对外部项目的依赖，主要是golang.org/x下的各个包，net、text、crypto...
```

Go项目结构最小布局
```go
// 在Go项目仓库根路径下
├── go.mod
├── LICENSE
├── xxx.go
├── yyy.go


// 在Go项目仓库根路径下
├── go.mod
├── LICENSE
├── package1
├── package1.go
├── package2
├── package2.go
...
```

[二进制可执行文件的Go项目结构](https://github.com/emorydu/iam)
```go
Project
├── LICNESE
├── Makefile
├── README.md
├── cmd/  // 主要存放所需要构建的可执行文件对应的main包的源文件
        // main包主要做命令行参数解析、资源初始化、日志设施初始化、数据库... 
				// 之后程序的控制权交给高级的执行控制对象
├──── app1
├────── main.go
├──── app2
├────── main.go
├── go.mod
├── go.sum
├── pkg/  // 存放项目自身要使用并且同样也是可执行文件对应main包要依赖的库文件
        // 可被外部项目引用
├──── lib1
├────── lib1.go
├──── lib2
├────── lib2.go
- [vendor]
```

[以只构建库为目的的Go项目结构](https://github.com/emorydu/log)
```go
Libarary
├── LICENSE
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── lib.go
├── lib1
├────── lib1.go
├── lib2
├────── lib2.go
```
