package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	val int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.val++
}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.val
}

func test() (res string) {
	res = "abccccc"
	return
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("a =", a.get())

	fmt.Println(test())
}