package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

// received-only channel: 这个chan只能收数据，不能送数据出去
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w)
	return w
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d recieved %c\n", id, n)
		w.wg.Done() // 通知已做完一个任务
	}
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//wg.Add(20)
	for i, worker := range workers {
		// 添加任务一定要在真实送入任务前，不然任务一旦送进去，可能还没Add到，就已经被worker Down了就会导致任务计数为负数panic
		wg.Add(1) // 每次添加一个任务
		worker.in <- 'a' + i
		//wg.Add(1) // 每次添加一个任务
	}

	for i, worker := range workers {
		wg.Add(1) // 每次添加一个任务
		worker.in <- 'A' + i
		//wg.Add(1) // 每次添加一个任务
	}

	wg.Wait()
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}
