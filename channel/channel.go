package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered Channel")
	bufferedChannel()
	fmt.Println("Channel close")
	closeChannel()
}

func chanDemo() {
	var c1 chan int // 这样定义的话 c == nil
	fmt.Println(c1 == nil)

	c := make(chan int)
	fmt.Println(len(c), cap(c))
	fmt.Println(c == nil)

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

// received-only channel: 这个chan只能收数据，不能送数据出去
func createWorker(id int) chan<- int {
	c := make(chan int)
	// 这里不能只for，不然死循环在for里，要异步goroutine起来
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	//for {
	//	n, open := <-c
	//	if !open {
	//		break
	//	} else {
	//		fmt.Printf("worker %d recieved %c\n", id, n)
	//	}
	//}
	// range也可以处理channel close
	for n := range c {
		fmt.Printf("worker %d recieved %c\n", id, n)
	}
}

func bufferedChannel() {
	c := make(chan int, 3) // buffer size = 3
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func closeChannel() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
