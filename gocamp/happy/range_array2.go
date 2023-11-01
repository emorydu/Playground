package main

import "fmt"

func arrayRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("arrayRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func pointerToArrayRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("pointerToArrayRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func sliceRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("sliceRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a[:] {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func sliceLenChangeRangeExpression() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	fmt.Println("sliceLenChangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}
		r = append(r, v)
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)

	/*
		原切片a在for range的循环过程中被附加了两个元素6和7，其len由5增加到7，
		但这对于r却没有产生影响。其原因就在于a的副本a'的内部表示中的len字段并没有改变，依旧是5，
		因此for range只会循环5次，也就只获取到a所对应的底层数组的前5个元素
	*/
}
