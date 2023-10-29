# Gopher09

- map原理

1. 何为map
```go
// map表示一组无序键值对
// map的key必须是可比较的(== !=)

var m map[string]int
m["key"] = 1

// 使用复合字面值创建map类型变量
// src/net/status.go
var statusText = map[int]string{
	StatusOK: "OK",
	StatusCreated: "Created",
	StatusAccepted: "Accepted",
    ...
}

// 使用make创建map类型变量
// src/net/client.go
icookies = make(map[string][]*Cookie)

// src/net/h2_bundle.go
http2commonLowerHeader = make(map[stirng]string, len(common))

// map与切片一样，在作为函数参数传递时不会有很大的性能损耗，并且在函数内部
// 对map变量的修改在外部可是可见的

var foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12
}

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m)
	foo(m)
	fmt.Println(m)
}
```

2. map基本操作
```go
// 插入: 前提（非nil）
// key存在使用新value覆盖旧值
m := make(map[K]V)
m[k1] = v1
m[k2] = v2
m[k3] = v3 

// 获取数据个数
m := map[string]int{
	"key1": 1,
	"key2": 2,
}
fmt.Println(len(m)) // 2
m["key3"] = 3
fmt.Println(len(m)) // 3

// 查找和数据读取 comma ok
_, ok := m["key"]
if !ok {
	// key 在map中不存在
}

m := map[string]int
m["key1"] = 1
m["key2"] = 2

v := m["key1"]
fmt.Println(v)  // 1
v = m["key3"]
fmt.Println(v)  // 'zero-value'

// comma-ok
m := map[string]int
v, ok := m["key"]
if !ok {
	
}
fmt.Println(v)

// 删除数据
// 即便要删除的数据在map中不存在，delete也不会导致panic
m := map[string]int {
	"key1": 1,
	"key2": 2,
}
fmt.Println(m)
delete(m, "key2")
fmt.Println(m)

// 遍历
// Go运行时在初始化map迭代器时对起始位置做了随机处理
// 因此千万不要依赖遍历map所得到的元素次序
func main() {
  m := map[int]int{
    1: 11, 
    2: 12, 
    3: 13, 
  }
  fmt.Println("{")
  for k, v := range m {
    fmt.Printf("[%d, %d] ", k, v)
  }
  fmt.Printf("}\n")
}
```

2. [map内部实现](https://weread.qq.com/web/reader/f343248072895ed9f34f408k0723244023c072b030ba601)

3. 总结
```go
// 不要依赖map的元素遍历顺序
// map不是线程安全的，不支持并发写
// 不要尝试获取map中元素（value）的地址
// 尽量使用cap参数创建map，以提升map平均访问性能，减少频繁扩容带来的不必要损耗
```