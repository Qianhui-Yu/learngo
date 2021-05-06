package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

// received-only channel: 这个chan只能收数据，不能送数据出去
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	// 这里不能只for，不然死循环在for里，要异步goroutine起来
	go doWork(id, w)
	return w
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d recieved %c\n", id, n)
		w.done <- true //用通信来共享内存
	}
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for _, worker := range workers {
		<-worker.done // 用来接收是否完成任务
		//因为worker.down是阻塞chan，这一步是block住的，所以只有这里读出了数据，才能输入大写字母打印的任务处理完成的信号
		//不然大写字母任务完成的down会被上一批down阻塞住导致deadlock
		//deadlock: 大写字母打印任务完成的done在等小写字母打印完成的down chan被读出，而done chan一直没人读
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	for _, worker := range workers {
		<-worker.done // 用来接收是否完成任务
		//在这里down不行，晚了
		//<-worker.done
	}
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}
