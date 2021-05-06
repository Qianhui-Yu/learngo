package main

import (
	"fmt"
	"math/rand"
	"time"
)

// received-only channel: 这个chan只能收数据，不能送数据出去
func createWorker(id int) chan<- int {
	c := make(chan int)
	// 这里不能只for，不然死循环在for里，要异步goroutine起来
	go doWork(id, c)
	return c
}

func doWork(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d recieved %d\n", id, n)
	}
}

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func main() {
	c1, c2 := generator(), generator()
	w := createWorker(0)
	var cache []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for  {
		var activeWorker chan<- int // = nil
		var activeValue int
		//有值才把activeWorker指向真正的worker，能接受n
		//没有值，activeWorker = nil，就算走到case activeWorker <- n也不会怎么样,nil chan可以接收数据
		if len(cache) > 0 {
			activeWorker = w
			activeValue = cache[0]
		}
		//如果产生数据快，消费数据慢，产生的数据没有存下来，会被覆盖掉
		select {
		//调度
		case n := <-c1:
			cache = append(cache, n)
		case n := <-c2:
			cache = append(cache, n)
		case activeWorker <- activeValue:
			cache = cache[1:]
		//两个请求之间超过800ms就超时
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Timeout!")
		//定时任务
		case <-tick:
			fmt.Println("Cache length =", len(cache))
		//总计时器
		case t := <-tm:
			fmt.Println("Time is up:", t)
			return
		//default:
		//	fmt.Println("Received nothing")
		}
	}
}
