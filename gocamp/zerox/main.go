package main

import (
	"fmt"
	"sync"
)

type User struct {
	sync.Mutex
	m map[int]int
}

func (u *User) Increment(wg *sync.WaitGroup) {
	u.Lock()
	defer wg.Done()
	defer u.Unlock()
	u.m[3]++

}

type Counter struct {
	sync.Mutex
	total int
}

func main() {
	var u = User{m: map[int]int{3: 0}}

	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go u.Increment(&wg)
	}

	wg.Wait()

	fmt.Println(u.m[3])

}
