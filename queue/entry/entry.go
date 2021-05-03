package main

import (
	"fmt"
	"learngo/queue"
)

func main() {
	q := queue.Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	fmt.Println(q)
	for q.IsEmpty() == false {
		fmt.Println(q.Pop())
	}
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())

}
