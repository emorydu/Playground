# Gopher14

- 使用if控制语句时遵循“快乐路径”原则
```go
func do() error {
  if errCondition1 {
    // ... err logic
    return err1 
  }
	
    // ... success logic
  if errCondition1 {
    // ... err logic
    return err1
  }
  // ...success logic
  return nil
}
// 当出现错误时，快速返回
// 成功逻辑不要嵌入if-else语句中
// “快乐路径”的执行逻辑在代码布局上始终靠左
// “快乐路径”的返回值一般在函数最后一行
```

- for避坑指南 

1. [小心迭代变量的重用](./range_array1.go)
2. [注意参与迭代的是range表达式的副本](./range_array2.go)
3. [其他range表达式类型的使用注意事项](./r_string.go)

- break避坑

Go语言规范中明确规定break语句（不接label的情况下）结束执行并跳出的是同一函数内break语句所在的最内层的for、switch或select的执行

- 尽量使用case表达式列表替代fallthrough

- 小结
1. 使用if遵循“快乐路径”原则
2. 小心for range的使用，明确真实参与循环的是range表达式的副本
3. 明确break和continue执行后的真实目的地
4. 使用fallthrough关键字前，考虑能否使用更简洁、清晰的case表达式列表替代