package main

import "sync"

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

func main() {

}
