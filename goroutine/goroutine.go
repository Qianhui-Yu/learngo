package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count int32
	for i := 0; i < 1000; i++ {
		go func(i int) {
			atomic.AddInt32(&count, 1)
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i) //此处一定要把i传进来，不然所有goroutine读的同一个i，
		// 当i=1500时，外部for循环其实结束了，但内部的func还是会用到这个i=1500，会产生意外的结果（race condition）
		// Data race: 当有多个协程在同时对一个数据进行读写时就会引发数据竞争
	}
	time.Sleep(3 * time.Millisecond)
	//runtime.Gosched()
	fmt.Printf("count = %d\n", count) // 此处和14行的写入count会产生data race，需要用channel来解决
}
